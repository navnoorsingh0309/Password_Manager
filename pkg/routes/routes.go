package routes

import (
	"jwt-app/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRotues = func(r *mux.Router) {
	r.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	r.HandleFunc("/signup", controllers.HandleSignUp).Methods("POST")
	r.HandleFunc("/getpasses", controllers.HandleGetPasswords).Methods("GET")
	r.HandleFunc("/newpass", controllers.HandleNewPassword).Methods("POST")
	r.HandleFunc("/editpass", controllers.HandleEditPassword).Methods("PUT")
	r.HandleFunc("/deletepass", controllers.HandleDeletePassword).Methods("DELETE")
}
