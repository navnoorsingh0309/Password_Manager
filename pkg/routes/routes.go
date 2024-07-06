package routes

import (
	"encoding/json"
	"fmt"

	"jwt-app/pkg/controllers"
	"jwt-app/pkg/database"
	"jwt-app/pkg/models"
	"log"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var RegisterUserRotues = func(r *mux.Router, store database.PostgresStore) {
	controllers.SetStore(store)
	// Login
	r.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	// New User
	r.HandleFunc("/signup", controllers.HandleSignUp).Methods("POST")
	// Getting Password
	r.HandleFunc("/getpasses", ProtectedWithJWT(controllers.HandleGetPasswords)).Methods("GET")
	// Adding New Password
	r.HandleFunc("/newpass", ProtectedWithJWT(controllers.HandleNewPassword)).Methods("POST")
	// Editing Password
	r.HandleFunc("/editpass", ProtectedWithJWT(controllers.HandleEditPassword)).Methods("PUT")
	// Deleteing Password
	r.HandleFunc("/deletepass", ProtectedWithJWT(controllers.HandleDeletePassword)).Methods("DELETE")

	// Temp Route
	r.HandleFunc("/getusers", controllers.HandleGetUsers).Methods("GET")
}

// Protected Routes
// Creating middleware
func ProtectedWithJWT(handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling middleware")

		tokenString := r.Header.Get("x-jwt-token")
		_, err := ValidateJWTToken(tokenString)
		if err != nil {
			// Writing json
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			jsonRes, err := json.Marshal(models.APIError{Error: "Invalid Token"})
			if err != nil {
				log.Fatal(err)
			}
			w.Write(jsonRes)
		}

		handleFunc(w, r)
	}
}

// Validating JWT Token
func ValidateJWTToken(tokenString string) (*jwt.Token, error) {
	// Parsing token
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// HMAC secret
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
