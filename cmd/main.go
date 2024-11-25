package main

import (
	"log"

	config "github.com/GDA35/ECOM/cfg"
	"github.com/GDA35/ECOM/cmd/api"
	"github.com/GDA35/ECOM/db"
)

func main() {
	cfg := &config.Config{
		DB: config.DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "your_user",
			Password: "your_password",
			Name:     "your_db_name",
			SSLMode:  "disable",
		},
	}

	postgresDB, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	defer postgresDB.Close()

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
