package controllers

import (
	"encoding/json"
	"jwt-app/pkg/database"
	"jwt-app/pkg/models"
	"log"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var store database.PostgresStore
var client database.MongoDBClient

// Returning Json
func WriteJson(w http.ResponseWriter, data any) {
	// Writing json
	w.Header().Add("Content-Type", "application/json")
	// For CORS policy
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	jsonRespose, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonRespose)
}

// Some parameters
func SetStore(s database.PostgresStore) {
	store = s
}
func SetMongoClient(c database.MongoDBClient) {
	client = c
}
func DecodeJWTToken(myToken string) (int, error) {
	var tokenClaim models.JWTTokenClaims

	token, err := jwt.ParseWithClaims(myToken, &tokenClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return -1, err
	}

	// Checking token validity
	if !token.Valid {
		return -1, err
	}

	return tokenClaim.Id, nil
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Login Request
	var req models.LoginUserReq

	// Get login information
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJson(w, models.LoginResponse{
			Error: "Error",
			Token: "",
			Id:    -1,
		})
		return
	}

	// Login to database
	currentUserId, jwt_token, err := store.Loginuser(&req)
	if err != nil {
		WriteJson(w, models.LoginResponse{
			Error: "Error",
			Token: "",
			Id:    -1,
		})
		return
	}

	// We can't use as chrome is banning third party cookies from 2025
	// Will store token in cookie
	// cookie := http.Cookie{
	// 	Name:     "jwt",
	// 	Value:    jwt_token,
	// 	MaxAge:   int(time.Now().Add(time.Minute * 30).Unix()),
	// 	HttpOnly: true,
	// }
	// // Setting Cookie
	// http.SetCookie(w, &cookie)

	WriteJson(w, models.LoginResponse{
		Error: "",
		Token: jwt_token,
		Id:    currentUserId,
	})
}

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	// Creating new request]
	createUserReq := new(models.CreateUserReq)
	// Checking for payload
	if err := json.NewDecoder(r.Body).Decode(createUserReq); err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}

	// Creating New User
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserReq.Password), bcrypt.DefaultCost)
	if err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}
	// Setting up user
	user, err := models.NewUser(createUserReq.Name, createUserReq.Email, encryptedPassword)
	if err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}

	// Adding user to database
	err = store.CreateUser(user, &client)
	if err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}

	// Writing json
	WriteJson(w, models.Message{Message: "success"})
}

func HandleNewPassword(w http.ResponseWriter, r *http.Request) {
	var getRequest models.PasswordModel
	// Decoding Id
	id, err := DecodeJWTToken(r.Header.Get("x-jwt-token"))
	if err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}

	// Decoding body
	if err := json.NewDecoder(r.Body).Decode(&getRequest); err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}

	// Creating new password for the user
	if err := client.NewPassword(id, &models.PasswordModel{
		Password: getRequest.Password,
		Entity:   getRequest.Entity,
	}); err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}
	WriteJson(w, models.Message{Message: "Success"})
}

func HandleGetPasswords(w http.ResponseWriter, r *http.Request) {
	// Decoding Id
	id, err := DecodeJWTToken(r.Header.Get("x-jwt-token"))
	if err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}

	// Getting all passwords wrt to Id
	passwords, err := client.GetPasswords(id)
	if err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}
	WriteJson(w, passwords)
}

func HandleEditPassword(w http.ResponseWriter, r *http.Request) {

}

func HandleDeletePassword(w http.ResponseWriter, r *http.Request) {

}

// Temp Route
func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := store.GetUsers()
	if err != nil {
		WriteJson(w, models.Message{Message: "Error"})
		return
	}

	// Writing json
	WriteJson(w, users)
}
