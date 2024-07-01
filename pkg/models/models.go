package models

import (
	"jwt-app/pkg/database"

	"go.mongodb.org/mongo-driver/mongo"
)

func Init() *mongo.Client {
	database.Connect()
	client := database.GetClient()
	return client

}

type User struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Username string `json:"Username"`
}
