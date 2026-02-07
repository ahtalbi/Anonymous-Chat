package main

import (
	"flag"
	"log"
	"net/http"

	"chat/internal/config"
	"chat/internal/handlers"
	"chat/internal/router"
	"chat/internal/server"
)

func main() {
	addr := flag.String("addr", ":8080", "this flag to specify the port")
	envirment  := flag.String("env", "dev", "dev or prod")
	flag.Parse()

	env, err := config.ParseEnv(*envirment)
	if err != nil {
		log.Fatal(err)
	}

	ser := server.New(&config.Config{
		Addr: *addr,
		Mux:  *http.NewServeMux(),
		Env:  env,
	})

	router := router.New(&ser.Config.Mux)
	router.Routes(map[string]func(w http.ResponseWriter, r *http.Request){
		"/": handlers.Home,
		"/api/auth": handlers.Login,
		"/api/session": handlers.CheckSession,
		"/api/ws": handlers.Ws,
	})

	ser.Listen()
}
