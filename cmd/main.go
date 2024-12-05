package main

import (
	"database/sql"
	"log"

	"github.com/GDA35/ECOM/cfg"
	"github.com/GDA35/ECOM/cmd/api"
	"github.com/GDA35/ECOM/db"
)

func main() {

	// Создание конфигурации
	config := cfg.NewConfig("localhost", "5432", "postgres", "Lbvjy5102006", "ecom", "disable")

	// Инициализация базы данных
	database, err := db.NewDatabase(config)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	defer func() {
		if err := database.Conn.Close(); err != nil {
			log.Println("Ошибка при закрытии соединения с базой данных:", err)
		}
	}()

	// Инициализация хранилища
	initStorage(database.Conn)

	// Запуск сервера API с передачей подключения к базе данных
	server := api.NewAPIServer(":8080", database.Conn)
	if err := server.Run(); err != nil {
		log.Fatal("Ошибка при запуске сервера API:", err)
	}

}

func initStorage(db *sql.DB) {
	// Проверка соединения с базой данных
	err := db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	log.Println("Успешное подключение к базе данных")
}
