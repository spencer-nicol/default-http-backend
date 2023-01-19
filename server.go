package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type server struct {
	ctx    context.Context
	router *http.ServeMux
}

func newServer(ctx context.Context) *server {
	s := &server{ctx, http.DefaultServeMux}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) ListenAndServe() {
	srv := http.Server{
		Addr:              ":8080",
		Handler:           s.router,
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		to, cancel := context.WithTimeout(s.ctx, 30*time.Second)
		defer cancel()
		if err := srv.Shutdown(to); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
