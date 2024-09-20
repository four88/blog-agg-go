package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/four88/blog-agg-go/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string
		Url  string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithErr(w, "Couldn't decode parameters", http.StatusInternalServerError)
		return
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithErr(w, "Couldn't create feed", http.StatusInternalServerError)
		return
	}

	responseWithJSON(w, databaseFeedToFeed(feed), http.StatusOK)
}

func (cfg *apiConfig) handleGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		responseWithErr(w, "Couldn't get feeds", http.StatusInternalServerError)
		return
	}
	responseWithJSON(w, databaseFeedsToFeeds(feeds), http.StatusOK)
}
