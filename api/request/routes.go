package request

import (
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1 *gin.RouterGroup) {
	var requestHandler Handler

	v1.GET("/request", requestHandler.Get, jwt.Verify)
	v1.POST("/request", requestHandler.Post, jwt.Verify)
	v1.PUT("/request/:id", requestHandler.Put, jwt.Verify)
	v1.DELETE("/request/:id", requestHandler.Delete, jwt.Verify)
}
