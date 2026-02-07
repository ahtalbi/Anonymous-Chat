package main

import (
	"flag"
	"log"
	"net/http"

	"chat/internal/config"
	"chat/internal/handlers"
	"chat/internal/realtime"
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

	clientManager := realtime.New()
	h := handlers.New(clientManager)

	router := router.New(&ser.Config.Mux)
	router.Routes(map[string]func(w http.ResponseWriter, r *http.Request){
		"/": h.Home,
		"/api/auth": h.Login,
		"/api/session": h.CheckSession,
		"/api/ws": h.Ws,
	})

	ser.Listen()
}
