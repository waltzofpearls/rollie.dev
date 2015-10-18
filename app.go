package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	config *Config
	router *mux.Router
}

func (a *App) NewApp(config *Config) *App {
	return &App{
		config: config,
		router: mux.NewRouter(),
	}
}

func (a *App) Run() {
	http.ListenAndServe(
		a.config.Address,
		a.router,
	)
}

func (a *App) Route(method, path string, h http.Handler) {
	a.router.Handle(path, h).Methods(method)
}
