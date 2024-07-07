package database

import (
	"context"
	"database/sql"
	"fmt"
	"jwt-app/pkg/models"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User database methods
type UserMethods interface {
	CreateTable() error
	CreateUser(*models.User) error
	LoginUser(*models.LoginUserReq) (*jwt.Token, error)
	// Temp Method
	GetUsers() ([]*models.User, error)
	GetUserById(int) (*models.User, error)
}

// Methods for particular user
type ParticularMethods interface {
	NewPassword(*models.PasswordModel) error
	DeletePassword(int) error
	UpdatePassword(*models.PasswordModel) error
}

// Postgres Storage
type PostgresStore struct {
	db *sql.DB
}

// MongoDB Client
type MongoDBClient struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
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

////////////////////////////////////////////////////////////////////////////////////////////////

// Initializing Postgres Storage
func NewPostgresStore() (*PostgresStore, error) {
	// Open Postgres service
	//connStr := "postgres://postgres:go_jwt@localhost:5432/postgres?sslmode=disable"
	connStr := "postgres://postgres:go_jwt@mypostgres:5432/postgres?sslmode=disable"
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

// Initializing MongoDB
func NewMongoDB() (*MongoDBClient, error) {
	//clientOptions := options.Client().ApplyURI("mongodb://mongoadmin:go_jwt@localhost:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.10")
	clientOptions := options.Client().ApplyURI("mongodb://mongoadmin:go_jwt@mymongo:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.10")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Pinging client
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	// Creating new database or getting existing one
	db := client.Database("user_passwords")

	return &MongoDBClient{
		client:     client,
		db:         db,
		collection: nil,
	}, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////

// Postgres Interface methods
func (s *PostgresStore) CreateUser(user *models.User, mongoClient *MongoDBClient) error {
	addUserQuery := `INSERT INTO USERS
					(name, email, encryptedPassword, create_at)
					VALUES ($1, $2, $3, $4)
					RETURNING id`
	// Adding new user to database
	var id int
	err := s.db.QueryRow(addUserQuery, user.Name, user.Email, user.EncryptedPassword, user.CreatedAt).Scan(&id)
	if err != nil {
		return err
	}

	// Create a New Colleciton
	if err := mongoClient.db.CreateCollection(context.TODO(), "user_"+strconv.Itoa(id)); err != nil {
		return err
	}

	fmt.Printf("Added User: %d\n", id)

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

// Scanning into users
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

////////////////////////////////////////////////////////////////////////////////////////////////

// Mongo DB New Collection for new signUp
func (client *MongoDBClient) NewCollection(user *models.User, name string) error {
	// Creating new collection with given user id
	if err := client.db.CreateCollection(context.TODO(), string(user.Id)); err != nil {
		return err
	}
	// Getting collection for user id
	client.collection = client.db.Collection(string(user.Id))
	return nil
}

// Mongo DB Database methods
func (client *MongoDBClient) NewPassword(passModel *models.PasswordModel) error {
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
		if err := res.Scan(&user.Id, &user.Name, &user.Email, &user.EncryptedPassword, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
