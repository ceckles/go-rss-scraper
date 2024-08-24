package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ceckles/go-rss-scraper/internal/auth"
	"github.com/ceckles/go-rss-scraper/internal/database"
	"github.com/google/uuid"
)

// handlerCreateUser creates a new user
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json: "name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithJSON(w, 400, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

// handlerGetUserByApiKey gets a user by api key
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error getting user: %v", err))
		return
	}
	usr, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't getting user: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(usr))

}
