package product

import (
	"database/sql"

	"github.com/GDA35/ECOM/types"
	"github.com/lib/pq"
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
	_, err := s.db.Exec(
		"INSERT INTO products (name, description, image, price, quantity) VALUES ($1, $2, $3, $4, $5)",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity,
	)
	if err != nil {
		return err
	}

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

func (s *Store) GetProductsByIDs(ids []int) ([]types.Product, error) {
	// Выполняем SQL-запрос для получения продуктов по идентификаторам
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ANY($1::integer[])", pq.Array(ids))

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

func (s *Store) UpdateProduct(product types.Product) error {
	// Выполняем SQL-запрос для обновления продукта
	_, err := s.db.Exec(
		"UPDATE products SET name = $1, description = $2, image = $3, price = $4, quantity = $5 WHERE id = $6",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity,
		product.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
