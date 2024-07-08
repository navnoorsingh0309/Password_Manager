# Password Manager

An easy-to-use prototype of a password manager, developed with HTML, CSS, and JavaScript for the frontend, and leveraging Go for the backend, designed to demonstrate basic functionalities such as user registration, login, and secure password storage.

## Features

- Authentication is using JWT Tokens rather than session tokens using the postgres database.
- Avoid 3rd party cookies during authentication.
- Handles all the CORS policies.
- Passwords are stored in mongodb database.
- Whole backend can dockerize using the docker compose file.
- As it is made in Go, so it scalabe and have quick compilation and execution

## Installation

### For backend

- Install docker in your system
- Run the folllowing commands:
  
  ` docker compose build `
  
  ` docker compose up `
  
### For fronteend

-  Just run the index.html file no need to install anything

## Backend details

- Backend has many routes like /login, /signup, /getpasses, /newpass, /getusers (just for testing), etc
- Out of these only /login and /signup can be access directly, others are protected routes and can be accessed using JWT Token.
  
### File description

- main.go is the main file for server that executes tasks from other files
- userControllers.go contains immeditate functions to routes.
- database.go contains functions with respect to database postgres and mongodb
- loadEnvVars.go loads the .env file
- models.go contains all the structures for different functions performed
- routes.go contains the initialization of routes we are using in backend