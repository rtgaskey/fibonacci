package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"main.go/pkg/service"
)

func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	current := service.GetCurrent()
	respondWithJSON(w, http.StatusOK, map[string]string{"message": current})
}

func NextHandler(w http.ResponseWriter, r *http.Request) {
	next := service.GetNext()

	respondWithJSON(w, http.StatusOK, map[string]string{"message": next})
}

func PreviousHandler(w http.ResponseWriter, r *http.Request) {
	previous := service.GetPrevious()

	respondWithJSON(w, http.StatusOK, map[string]string{"message": previous})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
