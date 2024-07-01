package models

import (
	"jwt-app/pkg/database"
)

func Init() {
	database.Connect()
}

type User struct {
	Name     string `json:"Name"`
	Username string `json:"Username"`
}

type MongoDetails struct {
	Email    string `json:"Email"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
}
