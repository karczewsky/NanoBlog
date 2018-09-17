package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"fmt"
)

func init() {
	fmt.Printf("Starting NanoBlog")
}

func main() {
	r := chi.NewRouter()

	// MIDDLEWARE
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(20 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world\n")
		fmt.Fprintf(w, "\nRequest:\n%#v", r)
	})

	http.ListenAndServe(":8000", r)
}
