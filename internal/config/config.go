package config

import (
	"fmt"
	"net/http"
)

type Env int

const (
	Dev Env = iota
	Prod
)

const (
	DOMAIN_NAME = "localhost"
)

type Config struct {
	Addr string
	Mux  http.ServeMux
	Env  Env
}

func ParseEnv(env string) (Env, error) {
	switch env {
	case "dev":
		return Dev, nil
	case "prod":
		return Prod, nil
	default:
		return 0, fmt.Errorf("invalid env: %s", env)
	}
}
