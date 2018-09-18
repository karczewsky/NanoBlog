package main

import "net/http"

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("test")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"alive\": true}"))
}
