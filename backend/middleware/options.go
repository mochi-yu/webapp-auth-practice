package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Options() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
