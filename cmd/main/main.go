package main

import (
	"fmt"
	"jwt-app/pkg/initializers"
	"jwt-app/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	// Setting up a new router for our routes to api
	mux := mux.NewRouter()
	// Registering routes
	routes.RegisterUserRotues(mux)
	http.Handle("/", mux)

	// Starting server
	fmt.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
