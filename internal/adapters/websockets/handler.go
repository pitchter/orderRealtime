package websockets

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
)

func (wsh *WebSocketHandler) HandleMessages() {
	for {
		msg := <-wsh.Broadcast
		message, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error marshalling message: %v", err)
			continue
		}
		for client := range wsh.Clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error writing message: %v", err)
				client.Close()
				delete(wsh.Clients, client)
			}
		}
	}
}
