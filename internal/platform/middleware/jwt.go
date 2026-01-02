package middleware

import (
	"fmt"
	"net/http"
	"strings"

	jwtutil "github.com/Hapaa16/janken/internal/platform/jwt"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		parts := strings.SplitN(auth, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format",
			})
			return
		}

		userID, err := jwtutil.Parse(parts[1])
		if err != nil {
			fmt.Println("here ?")
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		// Store userID for handlers & websocket
		c.Set("userID", userID)
		c.Next()
	}
}
