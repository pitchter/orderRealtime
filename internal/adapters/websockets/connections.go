package websockets

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

type WebSocketHandler struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan interface{}
}

func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan interface{}),
	}
}

func (wsh *WebSocketHandler) HandleConnections(c *websocket.Conn) {
	defer func() {
		delete(wsh.Clients, c)
		c.Close()
	}()
	wsh.Clients[c] = true

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
	}
}
