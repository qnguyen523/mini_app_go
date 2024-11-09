package main

import (
	// "context"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"googleauth/controllers"
	"googleauth/models"

	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	godotenv.Load(".env.local")
	// connect to database
	models.Connect()
	// set up routes
	controllers.SetupRoutes()
	// set up CRUD for Word model

	fmt.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// CRUD operations examples
	// controllers.CreateWord(models.Database)
}
