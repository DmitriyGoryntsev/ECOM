package api

import (
	"database/sql"
	"log"
	"net/http"

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

	usersHandler := user.NewHandler()
	usersHandler.RegisterRoutes(subrouter)

	log.Println("Server listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
