package main

import (
	"encoding/json"
	"fmt"
	"github.com/four88/blog-agg-go/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// create feed follow
func (cfg *apiConfig) handleFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithErr(w, "Couldn't decode parameters", http.StatusInternalServerError)
		return
	}

	// Debugging output
	fmt.Printf("UserID: %v\n", user.ID)
	fmt.Printf("FeedID: %v\n", params.FeedID)

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		fmt.Printf("Error creating feed follow: %v\n", err)
		responseWithErr(w, "Couldn't create feed follow", http.StatusInternalServerError)
		return
	}

	responseWithJSON(w, databaseFeedFollowToFeedFollow(feedFollow), http.StatusOK)
}

func (cfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt.Printf("UserID: %v\n", user.ID)
	feedFollows, err := cfg.DB.GetAllFeedFollowsByUser(r.Context(), user.ID)
	if err != nil {
		responseWithErr(w, "Couldn't get feed follows", http.StatusInternalServerError)
		return
	}
	response := []FeedFollow{}
	for _, feedFollow := range feedFollows {
		response = append(response, databaseFeedFollowToFeedFollow(feedFollow))
	}
	responseWithJSON(w, response, http.StatusOK)
}

// delete feed follow
func (cfg *apiConfig) handleFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithErr(w, "Couldn't decode parameters", http.StatusInternalServerError)
		return
	}
	feedFollow, err := cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		responseWithErr(w, "Couldn't delete feed follow", http.StatusInternalServerError)
		return
	}
	responseWithJSON(w, databaseFeedFollowToFeedFollow(feedFollow), http.StatusOK)
}
