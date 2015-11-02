package libs

import (
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (hf HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := hf(w, r); err != nil {
		log.Print("Internal error: " + err.Error())
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}
