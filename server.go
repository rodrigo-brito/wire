package main

import (
	"log"
	"net/http"
)

type Server interface {
	CityHandler(w http.ResponseWriter, r *http.Request)
}

type ServerOption func(*server)

type server struct {
	service WeatherService
}

func (s *server) CityHandler(w http.ResponseWriter, r *http.Request) {
	cityName := r.URL.Query().Get("q")
	result, err := s.service.ByCity(cityName)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.Write([]byte(result))
}

func NewServer(options ...ServerOption) Server {
	server := new(server)

	for _, option := range options {
		option(server)
	}

	return server
}

func WithWealthService(service WeatherService) ServerOption {
	return func(server *server) {
		server.service = service
	}
}
