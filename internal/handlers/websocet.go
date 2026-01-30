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
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Error in Upgrading the web socket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("ws read:", err)
			break
		}
		log.Println(messageType, string(message), err)
		err = conn.WriteMessage(websocket.TextMessage, []byte("message from backend : "+string(message)))
		if err != nil {
			log.Println("ws write:", err)
			break
		}
	}
}
