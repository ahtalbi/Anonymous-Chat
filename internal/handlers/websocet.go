package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Ws(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		log.Println(messageType, string(message), err)
	}
}
