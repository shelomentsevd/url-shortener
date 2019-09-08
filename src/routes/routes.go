package routes

import (
	"handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// HTML
	r.Handle("/", handler.NewPage()).Methods(http.MethodGet)
	// Create link
	r.Handle("/--", handler.NewCreate()).Methods(http.MethodGet)
	// Get link & redirect
	r.Handle("/{link: [a-zA-Z0-9]{5}}", handler.NewGet()).Methods(http.MethodGet)

	return r
}
