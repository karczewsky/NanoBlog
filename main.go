package main

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
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
