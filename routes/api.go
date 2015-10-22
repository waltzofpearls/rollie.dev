package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waltzofpearls/tetris.go/libs"
)

type Api struct {
	router *mux.Router
}

func (sr *Api) SetRouter(r *mux.Router) {
	sr.router = r
}

func (sr *Api) AttachHandlers() {
	sr.router.Handle("/", libs.HandlerFunc(notFoundHandler))
	sr.router.Handle("/projects", libs.HandlerFunc(projectsHandler))
	sr.router.Handle("/contributions", libs.HandlerFunc(contributionsHandler))
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
