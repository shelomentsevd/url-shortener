package routes

import (
	"handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/", handler.NewPage()).Methods(http.MethodGet)
	r.Handle("/--", handler.NewCreate()).Methods(http.MethodGet)
	r.Handle("/{link: [a-zA-Z0-9]{5}}", handler.NewGet()).Methods(http.MethodGet)

	return r
}