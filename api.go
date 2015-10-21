package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	subrouter *mux.Router
}

func (sr *Api) SetRouter(r *mux.Router) {
	sr.subrouter = r
}

func (sr *Api) AttachHandlers() {
	sr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
	sr.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
	})
	sr.HandleFunc("/contributions", func(w http.ResponseWriter, r *http.Request) {
	})
}
