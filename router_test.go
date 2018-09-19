package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRouter(t *testing.T) {
	r := getRouter()

	t.Run("Ping endpoint", func(t *testing.T) {
		rec := performRequest(r, "GET", "/ping")
		if rec.Code != http.StatusOK {
			t.Errorf("endpoint returned wrong status code: got %v want %v",
				rec.Code, http.StatusOK)
		}

		if expected := ("{\"alive\": true}"); string(rec.Body.Bytes()[:]) != expected {
			t.Errorf("endpoint returned wrong body: got %v want %v",
				rec.Body, expected)
		}
	})

}
