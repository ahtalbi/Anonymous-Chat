package models

import "github.com/gorilla/websocket"

type Client struct {
	Nickname  string `json:"nickname"`
	SessionId string
	Conn      *websocket.Conn
}
