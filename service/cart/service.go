package cart

import (
	"fmt"

	"github.com/GDA35/ECOM/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity: %d", item.Quantity)
		}

		productIDs[i] = item.ProductID
	}

	return productIDs, nil
}

func (h *Handler) createOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {

	productMap := make(map[int]types.Product)
	for _, p := range ps {
		productMap[p.ID] = p
	}

	// check if all products are available in the cart
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}
	//calculate total price
	totalPrice := calculateTotalPrice(items, productMap)

	//reduce quantity of products
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	//create the order
	orderID, err := h.store.CreateOrder(types.Order{UserID: userID, Total: totalPrice, Status: "pending", Address: "Some address"})
	if err != nil {
		return 0, 0, err
	}
	//create order items
	for _, item := range items {
		err := h.store.CreateOrderItem(types.OrderItem{OrderID: orderID, ProductID: item.ProductID, Quantity: item.Quantity, Price: productMap[item.ProductID].Price})
		if err != nil {
			return 0, 0, err
		}
	}

	return orderID, totalPrice, nil
}

func checkIfCartIsInStock(item []types.CartItem, productMap map[int]types.Product) error {
	if len(item) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, i := range item {
		product, ok := productMap[i.ProductID]
		if !ok {
			return fmt.Errorf("product %d not found", i.ProductID)
		}

		if product.Quantity < i.Quantity {
			return fmt.Errorf("product %d is out of stock", i.ProductID)
		}
	}

	return nil
}

func calculateTotalPrice(items []types.CartItem, productMap map[int]types.Product) float64 {
	var totalPrice float64
	for _, i := range items {
		product, ok := productMap[i.ProductID]
		if !ok {
			continue
		}
		totalPrice += product.Price * float64(i.Quantity)
	}
	return totalPrice
}
