package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config
var randomState string

func main() {
	// load environment variables
	godotenv.Load(".env.local")
	fmt.Println("ClientSecret", os.Getenv("ClientSecret"))
	fmt.Println("ClientID", os.Getenv("ClientID"))

	// Store the OAuth2 configuration globally
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:80/callback", // Set your redirect URL here
		ClientID:     os.Getenv("ClientID"),          // Replace with your Google Client ID
		ClientSecret: os.Getenv("ClientSecret"),      // Replace with your Google Client Secret
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	randomState = "random" // Use a secure random state in production

	// Handle the index and auth routes
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)
	fmt.Println("Server started at localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

// Display the login page with a Google login link
func handleIndex(w http.ResponseWriter, r *http.Request) {
	html := `<html><body><a href="/login">Log in with Google</a></body></html>`
	fmt.Fprintf(w, html)
}

// Redirect to Google's OAuth 2.0 login page
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Handle the OAuth 2.0 callback and get the user's info
func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// verify the state
	if r.FormValue("state") != randomState {
		http.Error(w, "State is invalid", http.StatusBadRequest)
		return
	}
	// Exchange authorization code for token
	token, err := googleOauthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, "Could not get auth token", http.StatusInternalServerError)
		return
	}
	// Fetch user information
	userInfo, err := getUserInfo(token)
	if err != nil {
		http.Error(w, "Could not get user info", http.StatusInternalServerError)
		return
	}
	// Display user information (you can store this info as needed)
	fmt.Fprintf(w, "User info: %v\n", userInfo)
	fmt.Println("token", token)
}

// Get user information from Google's UserInfo API
func getUserInfo(token *oauth2.Token) (userInfo map[string]interface{}, err error) {
	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode user info from response
	userInfo = make(map[string]interface{})
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
