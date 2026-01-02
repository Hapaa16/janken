package match

import (
	"github.com/Hapaa16/janken/internal/transport/websocket"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	CreateMatch(match *Match) error
	GetMatch(id string) (*Match, error)
	UpdateMatch(match *Match) error
	DeleteMatch(id string) error
}

type MatchRepository struct {
	db       *gorm.DB
	redis    *redis.Client
	hub      *websocket.Hub
	serverId string
}

func NewMatchRepository(db *gorm.DB, redis *redis.Client, hub *websocket.Hub, serverId string) *MatchRepository {
	return &MatchRepository{
		db:       db,
		redis:    redis,
		hub:      hub,
		serverId: serverId,
	}
}
