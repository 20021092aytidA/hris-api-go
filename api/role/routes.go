package role

import (
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1 *gin.RouterGroup) {
	var roleHandler Handler

	v1.GET("/role", roleHandler.Get, jwt.Verify)
	v1.POST("/role", roleHandler.Post, jwt.Verify)
	v1.PUT("/role/:id", roleHandler.Put, jwt.Verify)
	v1.DELETE("/role/:id", roleHandler.Delete, jwt.Verify)
}
