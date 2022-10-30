package models

import "github.com/gorilla/websocket"

type PublishRequest struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}
type Client struct {
	ID         string
	Connection *websocket.Conn
}
type Message struct {
	Action  string `json:"action"`
	Topic   string `json:"topic"`
	Message string `json:"message"`
}
