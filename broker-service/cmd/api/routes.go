package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routers() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connected
}
