package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GDA35/ECOM/cfg"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase(cfg *cfg.Config) (*Database, error) {
	conn, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Database{Conn: conn}, nil
}

func (db *Database) Close() {
	if err := db.Conn.Close(); err != nil {
		log.Println("failed to close database connection:", err)
	}
}
