package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testCase struct {
	path   string
	method string
	status int
	body   string
}

func TestHandleDefault(t *testing.T) {
	tc := []testCase{
		{"/", http.MethodGet, http.StatusNotFound, NotFound},
		{"/", http.MethodPost, http.StatusNotFound, NotFound},
		{"/", http.MethodPut, http.StatusNotFound, NotFound},
		{"/", http.MethodDelete, http.StatusNotFound, NotFound},
		{"/", http.MethodConnect, http.StatusNotFound, NotFound},
		{"/", http.MethodHead, http.StatusNotFound, NotFound},
		{"/", http.MethodOptions, http.StatusNotFound, NotFound},
		{"/", http.MethodPatch, http.StatusNotFound, NotFound},
		{"/", http.MethodTrace, http.StatusNotFound, NotFound},
		{"/test", http.MethodGet, http.StatusNotFound, NotFound},
		{"/test/", http.MethodGet, http.StatusNotFound, NotFound},
	}

	for idx, c := range tc {
		r := httptest.NewRequest(c.method, c.path, nil)
		w := httptest.NewRecorder()
		handleDefault().ServeHTTP(w, r)

		actual := w.Result().StatusCode
		if actual != c.status {
			t.Errorf("%d: expected %d; actual %d", idx, c.status, actual)
		}

		b, err := io.ReadAll(w.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) != c.body {
			t.Errorf("%d: expect %s; actual %s", idx, c.body, string(b))
		}
	}
}

func TestHandleHealthCheck(t *testing.T) {
	tc := []testCase{
		{"/healthz", http.MethodGet, http.StatusOK, Up},
		{"/healthz", http.MethodPost, http.StatusMethodNotAllowed, MethodNotAllowed},
		{"/healthz", http.MethodPut, http.StatusMethodNotAllowed, MethodNotAllowed},
		{"/healthz", http.MethodDelete, http.StatusMethodNotAllowed, MethodNotAllowed},
		{"/healthz", http.MethodConnect, http.StatusMethodNotAllowed, MethodNotAllowed},
		{"/healthz", http.MethodHead, http.StatusMethodNotAllowed, MethodNotAllowed},
		{"/healthz", http.MethodOptions, http.StatusMethodNotAllowed, MethodNotAllowed},
		{"/healthz", http.MethodPatch, http.StatusMethodNotAllowed, MethodNotAllowed},
		{"/healthz", http.MethodTrace, http.StatusMethodNotAllowed, MethodNotAllowed},
		{"/healthz/", http.MethodGet, http.StatusOK, Up},
	}

	for idx, c := range tc {
		r := httptest.NewRequest(c.method, c.path, nil)
		w := httptest.NewRecorder()
		handleHealthCheck().ServeHTTP(w, r)

		actual := w.Result().StatusCode
		if actual != c.status {
			t.Errorf("%d: expected %d; actual %d", idx, c.status, actual)
		}

		b, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Fatal(err)
		}

		if strings.TrimSpace(string(b)) != c.body {
			t.Errorf("%d: expect %s; actual %s", idx, c.body, string(b))
		}
	}
}
