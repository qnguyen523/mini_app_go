package handlers

import (
	"fmt"
	"main/models"

	// "main/utils"
	"net/http"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// register a new user
func Register(c *gin.Context) {
	fmt.Println("Received a registration request")
	// Parse request body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	// Check if the email already exists
	var existingUser models.User
	if err := models.Database.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}
	// Create new user
	user_input := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}
	// Insert user into database
	if err := models.Database.Create(&user_input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

var _ = test()

func test() string {
	return "test"
}

// get a user detail
func User(c *gin.Context) {
	fmt.Println("Request to get user")
	// get JWT token header
	// const bearerSchema = "Bearer "
	// authHeader := c.GetHeader("Authorization")
	// tokenString := authHeader[len(bearerSchema):]
	// fmt.Println("Auth header:", tokenString)
	// get JWT token from cookie
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// Pass jwt token with claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})
	// handle token parsing errors
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// get claims from token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse claims"})
		return
	}
	// check if token is expired
	if (*claims)["exp"].(float64) < float64(time.Now().Unix()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}
	// Get user id
	id, _ := strconv.Atoi((*claims)["sub"].(string))
	user := models.User{ID: uint(id)}
	// Query user from database using ID
	err = models.Database.Where("id =?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// Return user details as JSON response
	c.JSON(http.StatusOK, user)
}

// Function for logging in
func Login(c *gin.Context) {
	var user_params models.User
	// check user credentials and generate a jwt token
	if err := c.ShouldBindJSON(&user_params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse user data"})
		return
	}
	// Check if user exists
	var user models.User
	// check if credentials are valid
	models.Database.Where("email = ?", user_params.Email).First(&user)
	if user.ID == 0 {
		fmt.Println("User not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}
	// Compare passwords
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(user_params.Password))
	if err != nil {
		fmt.Println("Invalid Password:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}
	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(user.ID)),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})
	token, err := claims.SignedString([]byte("secretKey"))
	// token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Set JWT token in cookie
	c.SetCookie("jwt", token, 86400, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}
func Logout(c *gin.Context) {
	fmt.Println("Received a logout request")
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// Pass jwt token with claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})
	// handle token parsing errors
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
	// get claims from token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse claims"})
		return
	}
	// set expired time to 1 second ago
	(*claims)["exp"] = 0

	// Clear JWT token by setting an empty value and expired time in the cookie
	// Expired time is set to 1 second ago
	// c.SetCookie("jwt", "", -1, "/", "localhost", true, true)
	// Return success response indicating logout was successful
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}
