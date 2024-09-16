package main

import (
	"net/http"
)

func errorHandle(w http.ResponseWriter, r *http.Request) {
	responseWithErr(w, "Internal server error", http.StatusInternalServerError)
}
