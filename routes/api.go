package routes

import (
	"net/http"

	"github.com/waltzofpearls/tetris-go/libs"
)

type Api struct {
	github *libs.Github

	libs.Subrouter
}

func (sr *Api) AttachHandlers() {
	sr.github = libs.NewGithub(sr.Config)
	sr.Router.Handle("/", libs.HandlerFunc(sr.notFoundHandler)).Methods("GET")
	sr.Router.Handle("/projects", libs.HandlerFunc(sr.projectsHandler)).Methods("GET")
	sr.Router.Handle("/contributions", libs.HandlerFunc(sr.contributionsHandler)).Methods("GET")
}

func (sr *Api) notFoundHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (sr *Api) projectsHandler(w http.ResponseWriter, r *http.Request) error {
	sr.github.GetRepos()
	return nil
}

func (sr *Api) contributionsHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
