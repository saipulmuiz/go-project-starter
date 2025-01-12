package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPWhitelistMiddleware(allowedIPs []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		for _, ip := range allowedIPs {
			if clientIP == ip {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access Denied"})
	}
}
