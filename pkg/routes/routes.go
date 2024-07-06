package routes

import (
	"jwt-app/pkg/controllers"
	"jwt-app/pkg/database"

	"github.com/gorilla/mux"
)

var RegisterUserRotues = func(r *mux.Router, store database.PostgresStore) {
	controllers.SetStore(store)
	// Login
	r.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	// New User
	r.HandleFunc("/signup", controllers.HandleSignUp).Methods("POST")
	// Getting Password
	r.HandleFunc("/getpasses", controllers.HandleGetPasswords).Methods("GET")
	// Adding New Password
	r.HandleFunc("/newpass", controllers.HandleNewPassword).Methods("POST")
	// Editing Password
	r.HandleFunc("/editpass", controllers.HandleEditPassword).Methods("PUT")
	// Deleteing Password
	r.HandleFunc("/deletepass", controllers.HandleDeletePassword).Methods("DELETE")

	// Temp Route
	r.HandleFunc("/getusers", controllers.HandleGetUsers).Methods("GET")
}
