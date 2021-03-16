package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/noc-tech/todo/datastore"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	ds := datastore.NewDatastore([]*datastore.Todo{})
	r.Mount("/todos", todoResource{ds}.Routes())
	http.ListenAndServe(":3001", r)
}
