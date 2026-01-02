package websocket

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Event struct {
	UserID string          `json:"user_id"`
	Type   string          `json:"type"`
	Data   json.RawMessage `json:"data"`
}

func StartRedisSubscriber(
	ctx context.Context,
	rdb *redis.Client,
	hub *Hub,
	channel string,
) {
	sub := rdb.Subscribe(ctx, channel)

	go func() {
		for msg := range sub.Channel() {
			var evt Event
			if err := json.Unmarshal([]byte(msg.Payload), &evt); err != nil {
				continue
			}

			if client, ok := hub.Get(evt.UserID); ok {
				payload, _ := json.Marshal(evt)
				client.Send <- payload
			}
		}
	}()
}
