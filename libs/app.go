package libs

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	config   *Config
	router   *mux.Router
	template *Template
}

func NewApp(config *Config) *App {
	return &App{
		config:   config,
		router:   mux.NewRouter(),
		template: NewTemplate(""),
	}
}

func (a *App) Run() {
	http.ListenAndServe(
		a.config.Listen.Address,
		a.router,
	)
}

func (a *App) UseStaticRouter(path string) {
	fs := http.FileServer(http.Dir(path))
	a.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func (a *App) UseRouter(path string, sr Subroutable) {
	s := a.router.PathPrefix(path).Subrouter()
	sr.InitRouter(s, a.config, a.template)
	sr.AttachHandlers()
}
