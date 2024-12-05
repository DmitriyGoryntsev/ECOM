package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GDA35/ECOM/types"
	"github.com/GDA35/ECOM/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что тело запроса не пустое
	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("request body is empty"))
		return
	}
	defer r.Body.Close()

	// Создаем переменную для хранения продукта
	var product types.Product

	// Декодируем JSON из тела запроса в структуру продукта
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Создаем продукт в хранилище
	err := h.store.CreateProduct(product)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Возвращаем созданный продукт в формате JSON
	if err := utils.WriteJSON(w, http.StatusCreated, product); err != nil { // Передаем product
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	// Получаем список продуктов из хранилища
	products, err := h.store.GetAllProducts()
	if err != nil {
		// Если произошла ошибка, возвращаем статус 500 и сообщение об ошибке
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Кодируем список продуктов в JSON и отправляем его в ответе
	if err := utils.WriteJSON(w, http.StatusOK, products); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
