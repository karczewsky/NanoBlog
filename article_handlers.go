package main

import (
	"context"
	"errors"
	"log"
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

func createArticle(w http.ResponseWriter, r *http.Request) {
	data := &article{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, errBadRequest400)
		return
	}

	if data.Title == "" || data.Body == "" {
		render.Render(w, r, errBadRequest400)
		return
	}

	if err := data.insertIntoDB(); err != nil {
		log.Fatal(err)
		render.Render(w, r, errServerError500)
		return
	}

	render.JSON(w, r, data)
}

type article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (a *article) Bind(r *http.Request) error {
	if a == nil {
		return errors.New("article data was not provided")
	}

	a.ID = 0
	return nil
}

func (a *article) insertIntoDB() error {
	row := database.QueryRow("INSERT INTO articles(title, body) VALUES($1, $2) RETURNING id", a.Title, a.Body)
	err := row.Scan(&a.ID)

	if err != nil {
		return errors.New("error adding article to DB")
	}

	return nil
}

var (
	errBadRequest400 = &errResponse{
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessage:   "Bad Request",
	}
	errServerError500 = &errResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorMessage:   "Internal server error occured",
	}
	articles []article
)
