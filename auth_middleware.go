package main

import (
	"github.com/four88/blog-agg-go/internal/auth"
	"github.com/four88/blog-agg-go/internal/database"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handle authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithErr(w, "Couldn't find api key", http.StatusUnauthorized)
			return
		}
		user, err := cfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			responseWithErr(w, "Couldn't get user", http.StatusInternalServerError)
			return
		}
		handle(w, r, user)
	}
}
