package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/noc-tech/todo/datastore"
)

type todoResource struct {
	ds datastore.StorerI
}

// Routes creates a REST router for the todos resource
func (rs todoResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			next.ServeHTTP(w, r)
		})
	})
	r.Options("/", func(w http.ResponseWriter, r *http.Request) {})     // OPTIONS /todos - required for cors
	r.Get("/", rs.List)                                                 // GET /todos - read a list of todos
	r.Post("/", rs.Create)                                              // POST /todos - create a new todo and persist it
	r.Delete("/{id}", rs.Delete)                                        // DELETE /todos/{id} - delete a todo
	r.Options("/{id}", func(w http.ResponseWriter, r *http.Request) {}) // OPTIONS /todos/{id} - required for cors
	return r
}

func (rs todoResource) List(w http.ResponseWriter, r *http.Request) {
	todos := rs.ds.TodoList()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (rs todoResource) Create(w http.ResponseWriter, r *http.Request) {
	var o datastore.Todo
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "can't decode body"})
		return
	}
	_ = rs.ds.TodoCreate(&o)
	w.WriteHeader(http.StatusCreated)
}

func (rs todoResource) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := rs.ds.TodoDelete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "can't delete todo"})
		return
	}
	w.WriteHeader(http.StatusOK)
}
