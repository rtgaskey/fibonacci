package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"main.go/pkg/service"
)

func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	current := service.GetCurrent()
	// respondWithJSON(w, http.StatusOK, map[string]string{"Current": current})
	respondWithJSON(w, http.StatusOK, current)
}

func NextHandler(w http.ResponseWriter, r *http.Request) {
	next := service.GetNext()

	// respondWithJSON(w, http.StatusOK, map[string]string{"message": next})
	respondWithJSON(w, http.StatusOK, next)
}

func PreviousHandler(w http.ResponseWriter, r *http.Request) {
	previous := service.GetPrevious()

	// respondWithJSON(w, http.StatusOK, map[string]string{"message": previous})
	respondWithJSON(w, http.StatusOK, previous)
}

// func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
func respondWithJSON(w http.ResponseWriter, code int, payload int) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
