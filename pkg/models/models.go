package models

import (
	"time"
)

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

type APIError struct {
	Error string `json:"error"`
}

type JWTToken struct {
	Token string `json:"Token"`
}

func NewUser(Name, Email string, Password []byte) (*User, error) {
	return &User{
		Name:              Name,
		Email:             Email,
		EncryptedPassword: Password,
		CreatedAt:         time.Now().UTC(),
	}, nil
}
