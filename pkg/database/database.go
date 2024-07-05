package database

import (
	"database/sql"
	"jwt-app/pkg/models"

	_ "github.com/lib/pq"
)

type User interface {
	CreateUser(*models.User) error
	DeleteUser(int) error
	UpdateUser(*models.User) error
	GetAccountById(int) (*models.User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=pqgotest dbname=postgres password=go-jwt sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}
