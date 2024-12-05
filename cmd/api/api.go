package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/GDA35/ECOM/service/product"
	"github.com/GDA35/ECOM/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Создаем хранилище пользователей
	userStore := user.NewStore(s.db)
	usersHandler := user.NewHandler(userStore)
	usersHandler.RegisterRoutes(subrouter)

	// Создаем хранилище продуктов
	productStore := product.NewStore(s.db)
	productsHandler := product.NewHandler(productStore)
	productsHandler.RegisterRoutes(subrouter)

	log.Println("Server listening on", s.addr)

	// Обработка ошибки при запуске сервера
	if err := http.ListenAndServe(s.addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}

	return nil
}
