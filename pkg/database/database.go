package database

import (
	"database/sql"
	"fmt"
	"jwt-app/pkg/models"

	_ "github.com/lib/pq"
)

// User database methods
type UserMethods interface {
	CreateTable() error
	CreateUser(*models.User) error
	DeleteUser(int) error
	// Temp Method
	GetUsers() ([]*models.User, error)
	UpdateUser(*models.User) error
	GetUserById(int) (*models.User, error)
}

// Postgres Storage
type PostgresStore struct {
	db *sql.DB
}

// Initializing Postgres Storage
func NewPostgresStore() (*PostgresStore, error) {
	// Open Postgres service
	connStr := "user=postgres dbname=postgres password=go_jwt sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Pinging to server
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Returning Database
	return &PostgresStore{
		db: db,
	}, nil
}

// Creating table for sql
func (s *PostgresStore) CreateTable() error {
	// Creating table with given configs
	query := `CREATE TABLE IF NOT EXISTS USERS (
		id SERIAL  PRIMARY KEY,
		name VARCHAR(50),
		email VARCHAR(50),
		create_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

// Postgres Interface methods
func (s *PostgresStore) CreateUser(user *models.User) error {
	addUserQuery := `INSERT INTO USERS
					(name, email, create_at)
					VALUES ($1, $2, $3)`
	// Adding new user to database
	addUser, err := s.db.Query(addUserQuery, user.Name, user.Email, user.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", addUser)
	return nil
}

func (s *PostgresStore) DeleteUser(int) error {
	return nil
}

func (s *PostgresStore) UpdateUser(*models.User) error {
	return nil
}

func (s *PostgresStore) GetUserById(int) error {
	return nil
}

// Temp Function to test
func (s *PostgresStore) GetUsers() ([]*models.User, error) {
	res, err := s.db.Query("SELECT * FROM USERS")
	if err != nil {
		return nil, err
	}

	users := []*models.User{}
	for res.Next() {
		user := new(models.User)
		if err := res.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
