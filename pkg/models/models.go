package models

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Requests
type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email`
	EncryptedPassword []byte    `json:"encryptedPassword"`
	CreatedAt         time.Time `json:"createat"`
}

// Responses
type LoginResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
	Id    int    `json:"id"`
}
type Message struct {
	Message string `json:"message"`
}

type PasswordModel struct {
	Entity   string `json:"entity"`
	Password string `json:"password"`
}

// Getting body
type GetData struct {
	Id int `json:"id"`
}
type JWTTokenClaims struct {
	jwt.RegisteredClaims
	expiresAt int
	Id        int
}

func NewUser(Name, Email string, Password []byte) (*User, error) {
	return &User{
		Name:              Name,
		Email:             Email,
		EncryptedPassword: Password,
		CreatedAt:         time.Now().UTC(),
	}, nil
}
