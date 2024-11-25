package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GDA35/ECOM/config"
)

type PostgresDB struct {
	Conn *sql.DB
}

func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии соединения с базой данных: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при подключении к базе данных: %v", err)
	}

	log.Println("Успешное подключение к базе данных PostgreSQL")

	return &PostgresDB{Conn: db}, nil
}

func (p *PostgresDB) Close() error {
	return p.Conn.Close()
}
