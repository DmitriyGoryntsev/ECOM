package main

import (
	"log"
	"os"

	"github.com/GDA35/ECOM/cfg"
	"github.com/GDA35/ECOM/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
)

func main() {
	// TODO password switch
	config := cfg.NewConfig("localhost", "5432", "postgres", "????", "ecom", "disable")

	database, err := db.NewDatabase(config)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Создание нового драйвера для миграций
	driver, err := postgres.WithInstance(database.Conn, &postgres.Config{})
	if err != nil {
		log.Fatal("Не удалось создать драйвер миграции:", err)
	}

	// Создание нового экземпляра миграции с драйвером
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres", // Эта строка — имя для базы данных
		driver,
	)
	if err != nil {
		log.Fatal("Не удалось создать новый экземпляр миграции:", err)
	}

	// Выполнение миграций
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	log.Println("Миграция выполнена успешно")
}
