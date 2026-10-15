package request

import (
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1 *gin.RouterGroup) {
	var requestHandler Handler

	v1.GET("/request", jwt.Verify, requestHandler.Get)
	v1.POST("/request", jwt.Verify, requestHandler.Post)
	v1.PUT("/request/:id", jwt.Verify, requestHandler.Put)
	v1.DELETE("/request/:id", jwt.Verify, requestHandler.Delete)
}
