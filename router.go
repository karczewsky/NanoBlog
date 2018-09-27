package main

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func getRouter() *chi.Mux {
	r := chi.NewRouter()

	// MIDDLEWARE
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// HANDLERS
	r.Get("/ping", pingHandler)

	// REST
	r.Route("/api", func(r chi.Router) {
		r.Route("/articles", func(r chi.Router) {
			// r.Post("/", addArticleHandler)
			r.Route("/{articleID}", func(r chi.Router) {
				r.Use(articleCtxMiddleware)
				r.Get("/", getSingleArticleHandler)
			})
		})
	})

	return r
}

type key string

type errResponse struct {
	HTTPStatusCode int    `json:"-"`
	ErrorMessage   string `json:"message"`
}

func (e errResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"alive\": true}"))
}
