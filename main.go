package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"main.go/pkg/handler"

	_ "net/http/pprof"
)

var router *chi.Mux

func routers() *chi.Mux {
	router.Get("/current", handler.CurrentHandler)
	router.Get("/next", handler.NextHandler)
	router.Get("/previous", handler.PreviousHandler)

	return router
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
}

func main() {
	routers()

	port := "8080"

	// log.Printf("Starting server on http://0.0.0.0:%s", port)

	http.ListenAndServe("0.0.0.0:"+port, router)

	fmt.Println("GOT HERE")
	recover()
}
