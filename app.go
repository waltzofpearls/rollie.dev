package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/waltzofpearls/rollie.dev/libs"
)

type App struct {
	config   *libs.Config
	router   *mux.Router
	template *libs.Template
	http     *http.Server
	errChan  chan error
	quitChan chan os.Signal
}

func NewApp(config *libs.Config) *App {
	router := mux.NewRouter()
	readTimeout := 10 * time.Second
	writeTimeout := 20 * time.Second
	handler := handlers.LoggingHandler(os.Stdout, router)
	return &App{
		config:   config,
		router:   router,
		template: libs.NewTemplate(config.Template.Path),
		http: &http.Server{
			Addr:         config.Listen.HTTP,
			Handler:      handler,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
		errChan:  make(chan error),
		quitChan: make(chan os.Signal, 1),
	}
}

func (a *App) Run() {
	go func() {
		log.Println("listening and serving HTTP")
		if err := a.http.ListenAndServe(); err != nil {
			a.errChan <- err
		}
	}()

	signal.Notify(a.quitChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-a.errChan:
		log.Println("HTTP(S) server error:", err)
		return
	case sig := <-a.quitChan:
		log.Println(sig.String(), "caught, shutting down")
		return
	}
}

func (a *App) UseStaticRouter(path string) {
	fs := http.FileServer(http.Dir(path))
	a.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func (a *App) UseRouter(path string, sr libs.Subroutable) {
	s := a.router.PathPrefix(path).Subrouter()
	sr.InitRouter(s, a.config, a.template)
	sr.AttachHandlers()
}
