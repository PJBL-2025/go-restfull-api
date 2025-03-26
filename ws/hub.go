package ws

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Client struct {
	Conn    *websocket.Conn
	UserID  int
	AdminID int
}

type Hub struct {
	clients map[*Client]bool
	mu      sync.Mutex
}

var WebSocketHub = &Hub{
	clients: make(map[*Client]bool),
}

func (h *Hub) HandleConnections(c *websocket.Conn) {
	userID := c.Query("user_id")
	adminID := c.Query("admin_id")

	var uid, aid int
	fmt.Sscanf(userID, "%d", &uid)
	fmt.Sscanf(adminID, "%d", &aid)

	client := &Client{
		Conn:    c,
		UserID:  uid,
		AdminID: aid,
	}

	h.mu.Lock()
	h.clients[client] = true
	h.mu.Unlock()

	defer func() {
		h.mu.Lock()
		delete(h.clients, client)
		h.mu.Unlock()
		c.Close()
	}()

	for {
		var msg map[string]interface{}
		if err := c.ReadJSON(&msg); err != nil {
			break
		}
		h.BroadcastMessage(msg, uid, aid)
	}
}

func (h *Hub) BroadcastMessage(msg map[string]interface{}, userID, adminID int) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for client := range h.clients {
		if client.UserID == userID || client.AdminID == adminID {
			if err := client.Conn.WriteJSON(msg); err != nil {
				client.Conn.Close()
				delete(h.clients, client)
			}
		}
	}
}
