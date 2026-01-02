package websocket

import (
	"github.com/Hapaa16/janken/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, wsHandler *Handler) {
	r.GET("/ws", middleware.JWT(), wsHandler.Handle)
}
