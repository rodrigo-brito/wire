package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const apiKey = "9092670530cd215b30e2b787f0dce058"

func main() {
	service := NewWeatherService(WithAPIKey(apiKey))
	server := NewServer(WithWealthService(service))

	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)
	router.Get("/", server.CityHandler)

	fmt.Println("Server started at http://localhost:4000")
	http.ListenAndServe(":4000", router)
}
