package main

import "github.com/gorilla/mux"

type Api struct {
	subrouter *mux.Router
}

func (sr *Api) SetRouter(r *mux.Router) {
	sr.subrouter = r
}

func (sr *Api) AttachHandlers() {
}
