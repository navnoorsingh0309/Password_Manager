package database

import (
	"database/sql"
	"fmt"
	"jwt-app/pkg/models"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// User database methods
type UserMethods interface {
	CreateTable() error
	CreateUser(*models.User) error
	LoginUser(*models.LoginUserReq) (*jwt.Token, error)
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

// Creating JWT Token for user
func CreateJWT(user *models.User) (string, error) {
	// Claims
	claims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Minute * 30).Unix(), // 30 Minutes as it contains passwords
		"Email":     user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
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
		encryptedPassword VARCHAR(60),
		create_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

// Postgres Interface methods
func (s *PostgresStore) CreateUser(user *models.User) error {
	addUserQuery := `INSERT INTO USERS
					(name, email, encryptedPassword, create_at)
					VALUES ($1, $2, $3, $4)`
	// Adding new user to database
	addUser, err := s.db.Query(addUserQuery, user.Name, user.Email, user.EncryptedPassword, user.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", addUser)
	return nil
}

func (s *PostgresStore) Loginuser(login *models.LoginUserReq) (string, error) {
	// Get user by email from database
	user, err := s.GetUserByEmail(login.Email)
	if err != nil {
		return "", err
	}

	// Comparing Hashed Passwords
	if err := bcrypt.CompareHashAndPassword(user.EncryptedPassword, []byte(login.Password)); err != nil {
		return "", err
	}

	// Generating Token
	token, err := CreateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *PostgresStore) DeleteUser(int) error {
	return nil
}

func (s *PostgresStore) UpdateUser(*models.User) error {
	return nil
}

func (s *PostgresStore) GetUserByEmail(email string) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM USERS WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return ScanIntoUsers(rows)
	}

	return nil, fmt.Errorf("user not found")
}

func ScanIntoUsers(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)
	err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.EncryptedPassword,
		&user.CreatedAt,
	)

	return user, err
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
		if err := res.Scan(&user.Id, &user.Name, &user.Email, &user.EncryptedPassword, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
