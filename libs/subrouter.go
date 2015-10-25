package libs

import "github.com/gorilla/mux"

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
