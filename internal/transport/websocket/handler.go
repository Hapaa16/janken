package websocket

import (
	"context"
	"net/http"
	"time"

	rkeys "github.com/Hapaa16/janken/internal/infra/redis"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Handler struct {
	hub      *Hub
	rdb      *redis.Client
	serverID string
}

func NewHandler(hub *Hub, rdb *redis.Client, serverID string) *Handler {
	return &Handler{
		hub:      hub,
		rdb:      rdb,
		serverID: serverID,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	userID := c.GetString("userID")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	client := &Client{
		UserID: userID,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}

	h.hub.Add(client)

	ctx := context.Background()

	h.rdb.Set(ctx, rkeys.SocketOwner(userID), h.serverID, 60*time.Second)

	go h.writePump(client)
	h.readPump(client)
}

func (h *Handler) writePump(c *Client) {
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func (h *Handler) readPump(c *Client) {
	defer func() {
		h.hub.Remove(c.UserID)
		c.Conn.Close()
		h.rdb.Del(context.Background(), rkeys.SocketOwner(c.UserID))
	}()

	c.Conn.SetReadLimit(512)
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		if _, _, err := c.Conn.ReadMessage(); err != nil {
			break
		}
	}
}
