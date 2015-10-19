package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Subrouter interface {
	SetRouter(r *mux.Router)
	AttachHandlers()
}

type App struct {
	config *Config
	router *mux.Router
}

func NewApp(config *Config) *App {
	return &App{
		config: config,
		router: mux.NewRouter(),
	}
}

func (a *App) Run() {
	http.ListenAndServe(
		a.config.Listen.Address,
		a.router,
	)
}

func (a *App) Route(method, path string, h http.Handler) {
	a.router.Handle(path, h).Methods(method)
}

func (a *App) useStaticRouter(path string) {
	fs := http.FileServer(http.Dir(path))
	a.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func (a *App) useRouter(path string, sr Subrouter) {
	s := a.router.PathPrefix(path).Subrouter()
	sr.SetRouter(s)
	sr.AttachHandlers()
}
