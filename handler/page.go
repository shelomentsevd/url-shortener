package handler

import "net/http"

type Page struct {}

func NewPage() *Page {
	return &Page{}
}

func (handler *Page) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

}