package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Index struct {
	router *mux.Router
}

func (sr *Index) SetRouter(r *mux.Router) {
	sr.router = r
}

func (sr *Index) AttachHandlers() {
	sr.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := executeTemplate(w, "index", 200, nil)
		log.Printf("%s", err)
	})
}
