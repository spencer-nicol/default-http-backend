package main

import (
	"log"
	"net/http"
)

const (
	Up               = "{ \"status\": \"UP\" }"
	NotFound         = "404 Not Found"
	MethodNotAllowed = "405 Method Not Allowed"
)

func handleDefault() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(NotFound))
		if err != nil {
			log.Println(err)
		}
	}
}

func handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(Up))
			if err != nil {
				log.Println(err)
			}
		default:
			http.Error(w, MethodNotAllowed, http.StatusMethodNotAllowed)
		}
	}
}
