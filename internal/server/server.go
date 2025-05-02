package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateRouter(log *log.Logger) *Server {
	r := chi.NewRouter()
	r.Get("/", handlers.HandleMain)
	r.Post("/upload", handlers.HandleUploud)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: log,
		Server: srv,
	}

}
