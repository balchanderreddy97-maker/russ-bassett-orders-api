package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		apiKey := c.GetHeader("x-api-key")

		if apiKey != "russ123" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}