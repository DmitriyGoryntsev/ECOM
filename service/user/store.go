package user

import (
	"database/sql"
	"fmt"

	"github.com/GDA35/ECOM/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	row := s.db.QueryRow("SELECT * FROM users WHERE email = $1", email)

	u, err := ScanIntoUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return u, nil
}

func ScanIntoUser(row *sql.Row) (*types.User, error) {
	user := new(types.User)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	row := s.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	u, err := ScanIntoUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	if s.db == nil {
		return fmt.Errorf("database connection is nil")
	}

	_, err := s.db.Exec("INSERT INTO users (\"firstName\", \"lastName\", \"email\", \"password\", \"createdAt\") VALUES ($1, $2, $3, $4, $5)",
		user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
