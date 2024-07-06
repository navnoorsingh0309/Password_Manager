package controllers

import (
	"encoding/json"
	"jwt-app/pkg/database"
	"jwt-app/pkg/models"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var store database.PostgresStore

// Returning Json
func WriteJson(w http.ResponseWriter, data any) {
	// Writing json
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonRespose, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonRespose)
}

func SetStore(s database.PostgresStore) {
	store = s
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Login Request
	var req models.LoginUserReq

	// Get login information
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJson(w, models.APIError{Error: "Permission Denied"})
	}

	// Login to database
	jwt_token, err := store.Loginuser(&req)
	if err != nil {
		WriteJson(w, models.APIError{Error: err.Error()})
	}

	// Will store token in cookie
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    jwt_token,
		MaxAge:   int(time.Now().Add(time.Minute * 30).Unix()),
		HttpOnly: true,
	}
	// Setting Cookie
	http.SetCookie(w, &cookie)

	WriteJson(w, models.Message{Message: "success"})

}

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	// Creating new request
	createUserReq := new(models.CreateUserReq)
	// Checking for payload
	if err := json.NewDecoder(r.Body).Decode(createUserReq); err != nil {
		log.Fatal(err)
	}

	// Creating New User
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserReq.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	// Setting up user
	user, err := models.NewUser(createUserReq.Name, createUserReq.Email, encryptedPassword)
	if err != nil {
		log.Fatal(err)
	}

	// Adding user to database
	err = store.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	// Writing json
	WriteJson(w, user)
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
	WriteJson(w, users)
}
