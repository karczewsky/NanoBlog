package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var webServerAddr string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// DB
	database = connectToDb()
	initDB()

	// WEBAPP
	webServerAddr = os.Getenv("WEBSERVER_ADDR")

	fmt.Printf("Starting NanoBlog\n")
	fmt.Printf("Open browser at http://%v\n", webServerAddr)
}

func main() {
	defer database.Close()

	router := getRouter()
	s := &http.Server{
		Addr:           webServerAddr,
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
