package handler

import "net/http"

type Get struct {}

func NewGet() *Get {
	return &Get{}
}

func (handler *Get) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

}
