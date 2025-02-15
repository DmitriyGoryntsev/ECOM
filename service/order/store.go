package order

import (
	"database/sql"

	"github.com/GDA35/ECOM/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO orders (\"userId\", \"total\", \"status\", \"address\", \"createdAt\") VALUES ($1, $2, $3, $4, $5)", order.UserID, order.Total, order.Status, order.Address, order.CreatedAt)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (\"orderId\", \"productId\", \"quantity\", \"price\") VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
