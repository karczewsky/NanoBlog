package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func getRouter() *chi.Mux {
	r := chi.NewRouter()

	// MIDDLEWARE
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// HANDLERS
	r.Get("/ping", pingHandler)

	// REST
	r.Route("/api", func(r chi.Router) {
		r.Route("/articles", func(r chi.Router) {
			// r.Post("/", addArticleHandler)
			r.Get("/", pingHandler)
			r.Route("/{articleID}", func(r chi.Router) {
				// accessing URLParam example
				// articleID := chi.URLParam(r, "articleID")
			})
		})
	})

	return r
}
