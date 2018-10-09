package handler

import "net/http"

type Create struct {}

func NewCreate() *Create {
	return &Create{}
}

func (handler *Create) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

}