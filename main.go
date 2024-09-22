package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     "517381597560252",
		ClientSecret: "2213b158098fead084b9608652ed4ae9",
		RedirectURL:  "https://www.google.com/",
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}
)

func main() {
	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/callback", HandleCallback)
	http.ListenAndServe(":8080", nil)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="/login">Login with Facebook</a>`))
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	oauthConfig := oauthConfig
	token, err := oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://graph.facebook.com/v13.0/me?fields=name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Welcome, %s!", userInfo.Name)
}
