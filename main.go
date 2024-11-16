package main

import (
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

	fmt.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
