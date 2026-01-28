package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	addr  = flag.String("addr", ":3000", "cette argument pour prendre le port de serveur")
	users sync.Map
)

type User struct {
	Name string `json:"name"`
}

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/u", u)
	ser := http.Server{
		Addr:    *addr,
		Handler: mux,
	}
	fmt.Println("server is running on http://localhost" + *addr)
	log.Fatal(ser.ListenAndServe())
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func u(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		return
	}
	name := user.Name
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	_, loaded := users.LoadOrStore(name, conn)
	if loaded {
		w.Write([]byte(name + " is already taken use unique name"))
		w.WriteHeader(http.StatusOK)
		return
	}
	users.Range(func(key, value any) bool {
		name := key.(string)
		conn := value.(*websocket.Conn)
		fmt.Println(name, conn)
		return true
	})
	return
	for {
	}
}
