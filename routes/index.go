package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waltzofpearls/tetris.go/libs"
)

type Index struct {
	router *mux.Router
}

func (sr *Index) SetRouter(r *mux.Router) {
	sr.router = r
}

func (sr *Index) AttachHandlers() {
	sr.router.Handle("/", libs.HandlerFunc(indexHandler))
}

func indexHandler(w http.ResponseWriter, r *http.Request) error {
	return libs.ExecuteTemplate(w, "index", 200, nil)
}
