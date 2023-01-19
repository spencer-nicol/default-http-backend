package main

import (
	"context"
)

func main() {
	s := newServer(context.Background())
	s.ListenAndServe()
}
