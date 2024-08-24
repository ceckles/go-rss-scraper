package main

import (
	"fmt"
	"net/http"

	"github.com/ceckles/go-rss-scraper/internal/database"

	"github.com/ceckles/go-rss-scraper/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		handler(w, r, usr)
	}
}
