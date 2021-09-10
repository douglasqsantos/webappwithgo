package main

import (
	"net/http"

	"github.com/douglasqsantos/webappwithgo/config"
	"github.com/douglasqsantos/webappwithgo/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//Gracefully absorb panics and prints the stack trace
	mux.Use(middleware.Recoverer)
	// Logs the start and end of each request with the elapsed processing time
	mux.Use(middleware.Logger)
	// Generates the CSRF Token
	mux.Use(NoSurf)
	// Handle all the sessions.
	mux.Use(SessionLoad)

	// Routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
