package main

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Printf("Starting NanoBlog\n")
}

func main() {
	router := getRouter()

	http.ListenAndServe(":8000", router)
}
