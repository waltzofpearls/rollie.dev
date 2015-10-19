package main

import "github.com/gorilla/mux"

type Index struct {
	subrouter *mux.Router
}

func (sr *Index) SetRouter(r *mux.Router) {
	sr.subrouter = r
}

func (sr *Index) AttachHandlers() {
}
