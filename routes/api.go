package routes

import (
	"net/http"

	"github.com/waltzofpearls/tetris.go/libs"
)

type Api struct {
	libs.Subrouter
}

func (sr *Api) AttachHandlers() {
	sr.Router.Handle("/", libs.HandlerFunc(notFoundHandler))
	sr.Router.Handle("/projects", libs.HandlerFunc(projectsHandler))
	sr.Router.Handle("/contributions", libs.HandlerFunc(contributionsHandler))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func projectsHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func contributionsHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
