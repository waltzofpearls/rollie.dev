package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	router *mux.Router
}

func (sr *Api) SetRouter(r *mux.Router) {
	sr.router = r
}

func (sr *Api) AttachHandlers() {
	sr.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
	sr.router.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
	})
	sr.router.HandleFunc("/contributions", func(w http.ResponseWriter, r *http.Request) {
	})
}
