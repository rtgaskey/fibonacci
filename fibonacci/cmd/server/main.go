package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "net/http/pprof"
)

var router *chi.Mux

var currentIdx int = 0
var previous int = 0

func routers() *chi.Mux {
	router.Get("/current", CurrentHandler)
	router.Get("/next", NextHandler)
	router.Get("/previous", PreviousHandler)

	return router
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	fmt.Printf("%s \n", strconv.Itoa(fibonacci(6)))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	current := strconv.Itoa(fibonacci(currentIdx))
	respondWithJSON(w, http.StatusOK, map[string]string{"message": current})
}

func NextHandler(w http.ResponseWriter, r *http.Request) {
	next := strconv.Itoa(fibonacci(currentIdx + 1))
	currentIdx++

	respondWithJSON(w, http.StatusOK, map[string]string{"message": next})
}

func PreviousHandler(w http.ResponseWriter, r *http.Request) {
	previous := "0"

	if currentIdx > 0 {
		previous = strconv.Itoa(fibonacci(currentIdx - 1))

		currentIdx--
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": previous})
}

func main() {
	routers()

	port := "8080"

	log.Printf("Starting server on http://localhost:%s", port)

	http.ListenAndServe(":"+port, router)
}
