package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, auth *AuthHandler) {
	v1 := r.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/register", auth.Register)
		}
	}
}
