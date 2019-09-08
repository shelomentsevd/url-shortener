package handler

import (
	"controller"
	"errors"
	"net/http"
	"strings"
)

const (
	paramURL  = "url"
	paramJSON = "json"
)

var ErrTooShort = errors.New("URL is too short")

type Create struct {
	Controller controller.Controller
}

func NewCreate(controller controller.Controller) *Create {
	return &Create{
		Controller: controller,
	}
}

func (handler *Create) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	pJSON := strings.ToLower(query.Get(paramJSON)) == "true"
	pURL := query.Get(paramURL)

	key, err := handler.Controller.Save(pURL)
	if err != nil {
		// TODO: 500
	}

	// TODO: 200 OK
}
