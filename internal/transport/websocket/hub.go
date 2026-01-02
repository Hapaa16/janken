package websocket

import "sync"

type Hub struct {
	clients map[string]*Client
	mu      sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*Client),
	}
}

func (h *Hub) Add(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[client.UserID] = client
}

func (h *Hub) Remove(userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients, userID)
}

func (h *Hub) Get(userID string) (*Client, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	c, ok := h.clients[userID]
	return c, ok
}
