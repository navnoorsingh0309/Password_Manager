package controllers

import (
	"encoding/json"
	"jwt-app/pkg/models"
	"net/http"
)

var currentUser models.User

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func HandleSignUp(w http.ResponseWriter, r *http.Request) {

}

func HandleNewPassword(w http.ResponseWriter, r *http.Request) {

}

func HandleGetPasswords(w http.ResponseWriter, r *http.Request) {

}

func HandleEditPassword(w http.ResponseWriter, r *http.Request) {

}

func HandleDeletePassword(w http.ResponseWriter, r *http.Request) {

}
