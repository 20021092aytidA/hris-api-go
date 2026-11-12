package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() gin.HandlerFunc {
	var allowedIP = make(map[string]bool)
	allowedIP["http://127.0.0.1:3000"] = true
	allowedIP["http://localhost:3000"] = true
	allowedIP["http://127.0.0.1:5173"] = true
	allowedIP["http://localhost:5173"] = true

	return func(c *gin.Context) {
		ip := c.Request.Header.Get("Origin")
		if ip == "" || allowedIP[ip] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", ip)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}

			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "CORS Rejection, Not Allowed by CORS!",
			})
		}
	}
}
