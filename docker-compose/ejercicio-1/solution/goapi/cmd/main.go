package main

import (
	"net/http"

	goapi "github.com/carlostrejo2308/goapi/pkg"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Create a new router and use a logger middleware
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Main page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Welcome to this page! c:"))
	})

	// Handel the people end-points
	r.Mount("/people", goapi.PeopleResource{}.Routes())

	// Serve on the 7777 port
	http.ListenAndServe(":7777", r)
}
