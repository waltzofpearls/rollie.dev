package routes

import (
	"net/http"

	"github.com/waltzofpearls/rolli3.net/libs"
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
	return sr.JsonNotFoundHandler(w, r)
}

func (sr *Api) projectsHandler(w http.ResponseWriter, r *http.Request) error {
	repos, err := sr.github.GetRepos()
	if err != nil {
		return sr.JsonErrorHandler(w, r, err)
	}
	return sr.JsonResponseHandler(w, r, repos)
}

func (sr *Api) contributionsHandler(w http.ResponseWriter, r *http.Request) error {
	contribs, err := sr.github.GetContribs()
	if err != nil {
		return sr.JsonErrorHandler(w, r, err)
	}
	return sr.JsonResponseHandler(w, r, contribs)
}
