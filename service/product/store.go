package product

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

func (s *Store) GetAllProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := ScanIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.Product) error {
	// Выполняем SQL-запрос для вставки нового продукта
	result, err := s.db.Exec(
		"INSERT INTO products (name, description, image, price, quantity, \"createdAt\") VALUES ($1, $2, $3, $4, $5, $6)",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity,
		product.CreatedAt,
	)
	if err != nil {
		return err
	}

	// Получаем ID последней вставленной записи
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Устанавливаем ID для нового продукта
	product.ID = int(id)

	return nil
}

func ScanIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
