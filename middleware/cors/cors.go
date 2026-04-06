package cors

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	var allowedIP = make(map[string]bool)
	allowedIP[""] = true

	return func(c *gin.Context) {
		ip := c.Writer.Header().Get("Origin")
		if allowedIP[ip] {

			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}

			c.Next()
		} else {
			c.AbortWithStatusJSON(204, gin.H{
				"message": "CORS Rejection, Not Allowed by CORS!",
			})

			panic(errors.New("CORS ERROR REJECTION!"))
		}
	}
}
