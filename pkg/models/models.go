package models

import (
	"time"
)

// Authentication
type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
	Id    int    `json:"id"`
}

type User struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email`
	EncryptedPassword []byte    `json:"encryptedPassword"`
	CreatedAt         time.Time `json:"createat"`
}

type Message struct {
	Message string `json:"message"`
}

type PasswordModel struct {
	ID       int    `json:"id"`
	Entity   string `json:"entity"`
	Password string `json:"password"`
}

func NewUser(Name, Email string, Password []byte) (*User, error) {
	return &User{
		Name:              Name,
		Email:             Email,
		EncryptedPassword: Password,
		CreatedAt:         time.Now().UTC(),
	}, nil
}
