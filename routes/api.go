package routes

import (
	"encoding/json"
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
	repos, err := sr.github.GetRepos()
	if err != nil {
		return sr.JsonErrorHandler(w, r, err)
	}
	if err := json.NewEncoder(w).Encode(repos); err != nil {
		return sr.JsonErrorHandler(w, r, err)
	}
	return nil
}

func (sr *Api) contributionsHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
