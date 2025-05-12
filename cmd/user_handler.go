package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"fmt"
	"github.com/Auxesia23/todo_list/internal/models"
	"github.com/Auxesia23/todo_list/internal/utils"
)


func (app *application) RegisterUser (w http.ResponseWriter, r *http.Request){
	//Decode request body dan bind ke variabel
	var userInput models.UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//Verifikasi apakah email dan password sudah valid
	err = utils.VerifyPasswordAndEmail(userInput.Password, userInput.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//Hash password sebelum di simpan ke db
	hashedPass, _ := utils.HashPassword(userInput.Password)
	user := models.User {
		Email: userInput.Email,
		Username: userInput.Username,
		Password: hashedPass,
	}
	
	//Menyimpan data user baru ke db
	userResponse, err := app.User.Create(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}

func (app *application) UserLogin(w http.ResponseWriter, r *http.Request){
	//Decode request body dan bind ke variabel
	var loginInput models.UserLogin
	err := json.NewDecoder(r.Body).Decode(&loginInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//Verifikasi credensial user dan membuat token jwt
	token, err := app.User.Login(context.Background(), loginInput.Email, loginInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (app *application) UserInfo (w http.ResponseWriter, r *http.Request){
	userEmail := r.Context().Value("userEmail").(string)
	user, err := app.User.Get(context.Background(), userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (app *application) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	redirectURI := os.Getenv("GOOGLE_REDIRECT_URI")
	scope := "email profile"
	authURL := "https://accounts.google.com/o/oauth2/v2/auth"
	
	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("redirect_uri", redirectURI)
	params.Add("response_type", "code")
	params.Add("scope", scope)
	params.Add("access_type", "offline")
	params.Add("prompt", "consent")
	fullURL := fmt.Sprintf("%s?%s", authURL, params.Encode())
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"url": fullURL})
}


func (app *application) GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	token, err := app.User.GoogleLogin(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}