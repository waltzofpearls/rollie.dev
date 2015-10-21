package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Index struct {
	subrouter *mux.Router
}

func (sr *Index) SetRouter(r *mux.Router) {
	sr.subrouter = r
}

func (sr *Index) AttachHandlers() {
	sr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
}
