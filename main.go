package main

import (
	// "net/http"
	"github.com/gin-gonic/gin"

	// "main/controllers"
	"main/models"

	"github.com/joho/godotenv"
	"main/handlers"
	"main/middleware"
)

func main() {
	r := gin.Default()
	// Public routes (do not require authentication)
	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/register", handlers.Register)
		publicRoutes.GET("/user", handlers.User)
		publicRoutes.GET("/logout", handlers.Logout)
	}
	// Protected routes (require authentication)
	protectedRoutes := r.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{

	}
	godotenv.Load(".env.local")
	models.Connect()
	r.Run(":8080")
	

	// godotenv.Load(".env.local")
	// handler := controllers.New()
	// server := &http.Server{
	// 	Addr:    "0.0.0.0:8008",
	// 	Handler: handler,
	// }
	// models.Connect()
	// server.ListenAndServe()
}
