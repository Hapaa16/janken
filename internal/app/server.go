package app

import (
	"context"
	"log"

	"github.com/Hapaa16/janken/internal/config"
	"github.com/Hapaa16/janken/internal/infra/db"
	"github.com/Hapaa16/janken/internal/infra/db/models"
	"github.com/Hapaa16/janken/internal/infra/redis"
	"github.com/gin-gonic/gin"

	// domain
	authDomain "github.com/Hapaa16/janken/internal/domain/auth"

	rkeys "github.com/Hapaa16/janken/internal/infra/redis"
	authHTTP "github.com/Hapaa16/janken/internal/transport/http"
	"github.com/Hapaa16/janken/internal/transport/websocket"
)

func Run() {
	// Load config
	cfg := config.Load()

	// Gin mode
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Gin engine
	router := gin.New()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	gormDB := db.New(cfg.DB)

	if err := gormDB.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatal(err)
	}

	rdb := redis.New(cfg.Redis)

	hub := websocket.NewHub()

	wsHandler := websocket.NewHandler(
		hub,
		rdb,
		cfg.ServerId,
	)

	websocket.StartRedisSubscriber(
		context.Background(),
		rdb,
		hub,
		rkeys.WSEventsChannel(),
	)

	authRepo := authDomain.NewAuthRepository(gormDB)

	authService := authDomain.NewService(authRepo)

	authHandler := authHTTP.NewAuthHandler(authService)

	authHTTP.RegisterRoutes(router, authHandler)

	websocket.RegisterRoutes(router, wsHandler)

	log.Printf("Server running on :%s", cfg.Port)
	log.Fatal(router.Run(":" + cfg.Port))
}
