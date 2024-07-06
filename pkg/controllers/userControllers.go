package controllers

import (
	"encoding/json"
	"jwt-app/pkg/database"
	"jwt-app/pkg/models"
	"log"
	"net/http"
)

var store database.PostgresStore

func SetStore(s database.PostgresStore) {
	store = s
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	// Creating new request
	createUserReq := new(models.CreateUserReq)
	// Checking for payload
	if err := json.NewDecoder(r.Body).Decode(createUserReq); err != nil {
		log.Fatal(err)
	}

	// Creating New User
	user := models.NewUser(createUserReq.Name, createUserReq.Email)
	if err := store.CreateUser(user); err != nil {
		log.Fatal(err)
	}

	// Writing json
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonRespose, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonRespose)
}

func HandleNewPassword(w http.ResponseWriter, r *http.Request) {

}

func HandleGetPasswords(w http.ResponseWriter, r *http.Request) {

}

func HandleEditPassword(w http.ResponseWriter, r *http.Request) {

}

func HandleDeletePassword(w http.ResponseWriter, r *http.Request) {

}

// Temp Route
func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := store.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	// Writing json
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonRespose, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonRespose)

}
