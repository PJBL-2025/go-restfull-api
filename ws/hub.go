package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

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

func (h *Hub) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Gagal upgrade ke WebSocket:", err)
		return
	}

	userID := r.URL.Query().Get("user_id")
	adminID := r.URL.Query().Get("admin_id")

	var uid, aid int
	fmt.Sscanf(userID, "%d", &uid)
	fmt.Sscanf(adminID, "%d", &aid)

	client := &Client{
		Conn:    conn,
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
		conn.Close()
	}()

	for {
		var msg map[string]interface{}
		err := conn.ReadJSON(&msg)
		if err != nil {
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
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				client.Conn.Close()
				delete(h.clients, client)
			}
		}
	}
}
