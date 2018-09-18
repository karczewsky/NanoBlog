package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"fmt"
)

func init() {
	fmt.Printf("Starting NanoBlog\n")
}

func main() {
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

	http.ListenAndServe(":8000", r)
}
