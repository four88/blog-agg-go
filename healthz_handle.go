package main

import (
	"net/http"
)

type responseHealth struct {
	Status string `json:"status"`
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	var res = responseHealth{}
	res.Status = "ok"
	responseWithJSON(w, res, http.StatusOK)
}
