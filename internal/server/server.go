package server

import (
	"log"
	"net/http"

	"chat/internal/config"
)

type ServerLocal struct {
	Config *config.Config
}

func New(cnf *config.Config) *ServerLocal {
	if cnf.Env == config.Dev {
		log.Println("debug enabled")
	}

	if cnf.Env == config.Prod {
		log.SetFlags(0)
	}

	return &ServerLocal{
		Config: cnf,
	}
}

func (s *ServerLocal) Listen() {
	server := http.Server{
		Addr: s.Config.Addr,
		Handler: &s.Config.Mux,
	}
	log.Println("http://"+config.DOMAIN_NAME+s.Config.Addr)
	log.Fatal(server.ListenAndServe())
}
