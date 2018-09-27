package main

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
	// godotenv.Load(".env")
	// os.Getenv("Password")
	articles = []article{
		article{
			ID:     1,
			Title:  "New article",
			Author: "Foo",
			Body:   "Bar",
		},
		article{
			ID:     2,
			Title:  "Fresh article",
			Author: "Foo",
			Body:   "Bar",
		},
	}

	fmt.Printf("Starting NanoBlog\n")
}

func main() {
	router := getRouter()

	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
