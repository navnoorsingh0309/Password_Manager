package routes

import (
	"jwt-app/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRotues = func(r *mux.Router) {
	r.HandleFunc("/getUsers", controllers.GetUsers).Methods("GET")
}
