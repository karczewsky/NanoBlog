package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func articleCtxMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID, err := strconv.Atoi(chi.URLParam(r, "articleID"))

		if err != nil {
			render.Render(w, r, errBadRequest400)
			return
		}

		ctx := context.WithValue(r.Context(), key("articleID"), articleID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getSingleArticleHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key("articleID"))
	art := article{}
	for _, a := range articles {
		if a.ID == id {
			art = a
			break
		}
	}

	if art == (article{}) {
		render.Render(w, r, errBadRequest400)
		return
	}

	render.JSON(w, r, art)
}

type article struct {
	ID    int
	Title string
	Body  string
}

var (
	errBadRequest400 = &errResponse{HTTPStatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request"}
	articles         []article
)
