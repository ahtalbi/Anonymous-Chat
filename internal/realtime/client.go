package realtime

import "github.com/gorilla/websocket"

type Client struct {
	Nickname string
	Conn     *websocket.Conn
}
