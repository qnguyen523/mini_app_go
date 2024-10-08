package utils

import (
	"fmt"
	// "time"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secretpassword")

// GenerateToken function generates a jwt token with the user id as part of the claims.
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token is valid for 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// verify a jwt token
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		// if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
		// 	return nil, fmt.Errorf("Invalid signing method")
		// }
		return secretKey, nil
	})
	// Check for errors
	if err != nil {
		return nil, err
	}
	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")

}
