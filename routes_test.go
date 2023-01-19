package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	tc := []struct {
		path   string
		method string
		status int
		body   string
	}{
		{"/", http.MethodGet, http.StatusNotFound, NotFound},
		{"/test", http.MethodGet, http.StatusNotFound, NotFound},
		{"/healthz", http.MethodGet, http.StatusOK, Up},
		{"/healthz", http.MethodPost, http.StatusMethodNotAllowed, MethodNotAllowed},
	}

	srv := server{context.Background(), http.DefaultServeMux}
	srv.routes()

	for idx, c := range tc {
		r := httptest.NewRequest(c.method, c.path, nil)
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, r)

		actual := w.Result().StatusCode
		if actual != c.status {
			t.Errorf("%d: expected %d; actual %d", idx, c.status, actual)
		}

		b, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Fatal(err)
		}

		if strings.TrimSpace(string(b)) != c.body {
			t.Errorf("%d: expected %s; actual %s", idx, c.body, string(b))
		}
	}
}
