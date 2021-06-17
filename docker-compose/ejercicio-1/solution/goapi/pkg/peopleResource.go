package goapi

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CarlosTrejo2308/peopleApiResource/abort"
	"github.com/go-chi/chi/v5"
)

type PeopleResource struct{}

// Routes returns a chi.Router that handles
// both the people endpoint and he people/{id} endpoints
func (rs PeopleResource) Routes() chi.Router {
	// Get new router
	r := chi.NewRouter()

	// Handle the main 'people' endpoint
	r.Get("/", rs.List)

	// Hanlde the /{id} endpoints
	r.Route("/{id}", func(r chi.Router) {
		// Get the id value and handle the ids endpoints
		r.Use(postCtx)
		r.Get("/", rs.Get)
	})

	return r
}

// List recives a http response writer and a http request
// and writes to the response writer of all the persons documents
// with a json format.
func (rs PeopleResource) List(w http.ResponseWriter, r *http.Request) {
	// Set the header
	w.Header().Set("Content-Type", "application/json")

	// Get the persons array of all the documents
	resp := GetAll()

	// Write the response to the responsewriter
	err := json.NewEncoder(w).Encode(resp)
	abort.AbortOnError(err)
}

// postCtx recives a handler and return another handle
// with the value of the id foun in the URL parameter as a context
func postCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the value of the id found in the url parameter
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		// Add the context to the given handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Get recives a responsewriter and http request and
// and writes to the response writer all the information
// of a person given an id, with a json format.
func (rs PeopleResource) Get(w http.ResponseWriter, r *http.Request) {
	// Get the id
	id := r.Context().Value("id").(string)

	// Convert to int
	idint, _ := strconv.Atoi(id)

	// Get the person information
	resp := GetById(idint)

	// Set the header
	w.Header().Set("Content-Type", "application/json")

	// Write the response to the responsewriter
	err := json.NewEncoder(w).Encode(resp)
	abort.AbortOnError(err)
}
