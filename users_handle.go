package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/four88/blog-agg-go/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithErr(w, "Couldn't decode parameters", http.StatusInternalServerError)
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithErr(w, "Couldn't create user", http.StatusInternalServerError)
		return
	}

	responseWithJSON(w, databaseUserToUser(user), http.StatusOK)
}

func (cfg *apiConfig) handlerUsersGetInfo(w http.ResponseWriter, r *http.Request, user database.User) {

	response := databaseUserToUser(user)
	responseWithJSON(w, response, http.StatusOK)
}
