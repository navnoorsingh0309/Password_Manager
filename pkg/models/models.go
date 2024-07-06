package models

import (
	"time"
)

type CreateUserReq struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type User struct {
	Id        int       `json:"Id"`
	Name      string    `json:"name"`
	Email     string    `json:"email`
	CreatedAt time.Time `json:"createat"`
}

func NewUser(Name, Email string) *User {
	return &User{
		Name:      Name,
		Email:     Email,
		CreatedAt: time.Now().UTC(),
	}
}
