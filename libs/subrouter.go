package libs

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Subroutable interface {
	InitRouter(*mux.Router, *Config)
	AttachHandlers()
}

type Subrouter struct {
	Router *mux.Router
	Config *Config
}

func (sr *Subrouter) InitRouter(r *mux.Router, c *Config) {
	sr.Router = r
	sr.Config = c
}

func (sr *Subrouter) RedirectHandler(w http.ResponseWriter, r *http.Request, url string) error {
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}

func (sr *Subrouter) HtmlResponseHandler(w http.ResponseWriter, r *http.Request, template string, data interface{}) error {
	buf, err := ExecuteTemplate(template, data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
	return nil
}

func (sr *Subrouter) JsonResponseHandler(w http.ResponseWriter, r *http.Request, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return sr.JsonErrorHandler(w, r, err)
	}
	return nil
}

func (sr *Subrouter) JsonNotFoundHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	return nil
}

func (sr *Subrouter) JsonErrorHandler(w http.ResponseWriter, r *http.Request, err error) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(JsonError{err.Error()}); err != nil {
		panic(err)
	}
	return nil
}
