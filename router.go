package main

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func getRouter() *chi.Mux {
	r := chi.NewRouter()

	// MIDDLEWARE
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(20 * time.Second))

	// HANDLERS
	r.Get("/ping", pingHandler)
	r.Get("/panic", panicHandler)

	// REST
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", pingHandler)

		r.Route("/{articleID}", func(r chi.Router) {
			// accessing URLParam example
			// articleID := chi.URLParam(r, "articleID")
		})
	})

	return r
}
