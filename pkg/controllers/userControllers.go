package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"jwt-app/pkg/database"
	"jwt-app/pkg/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CountUsers(w http.ResponseWriter, r *http.Request) {
	// Getting client
	client := database.GetClient()
	// Getting collection
	coll := client.Database("Users").Collection("user_details")
	// Setting options for counting
	opts := options.Count().SetHint("_id_")
	// Gettting count
	count, err := coll.CountDocuments(context.TODO(), bson.D{}, opts)
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(w, "%d", count)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Getting client
	client := database.GetClient()
	// Getting collection
	coll := client.Database("Users").Collection("user_details")
	// Setting options for find
	opts := options.Find()
	// Getting cursor for
	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		panic(err)
	}

	// Slice of Users
	Users := make(map[string]models.User)

	// Iterate over all docs
	for cursor.Next(context.TODO()) {
		// Create a value into which the single document can be decoded
		var user models.MongoDetails
		err := cursor.Decode(&user)
		if err != nil {
			panic(err)
		}

		Users[user.Email] = models.User{
			Name:     user.Name,
			Username: user.Username,
		}
	}

	if err := cursor.Err(); err != nil {
		panic(err)
	}

	// Marshelling to json
	jsonUsers, err := json.Marshal(Users)
	if err != nil {
		panic(err)
	}

	// Settin json Content
	w.Header().Set("Content-Type", "application/json")

	// Close the cursor
	cursor.Close(context.TODO())
	fmt.Fprintf(w, "%s", jsonUsers)
}
