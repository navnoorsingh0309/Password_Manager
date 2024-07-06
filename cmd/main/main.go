package main

import (
	"fmt"
	"jwt-app/pkg/database"
	"jwt-app/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Postgres
	store, err := database.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Database: %+v\n", store)

	// Setting up SQL Table
	if err := store.CreateTable(); err != nil {
		log.Fatal(err)
	}

	// Setting up a new router for our routes to api
	mux := mux.NewRouter()
	// Registering routes
	routes.RegisterUserRotues(mux, *store)
	http.Handle("/", mux)

	// Starting server
	fmt.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
