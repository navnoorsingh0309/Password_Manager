package controllers

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
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
