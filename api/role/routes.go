package role

import (
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1 *gin.RouterGroup) {
	var roleHandler Handler

	v1.GET("/role", jwt.Verify, roleHandler.Get)
	v1.POST("/role", jwt.Verify, roleHandler.Post)
	v1.PUT("/role/:id", jwt.Verify, roleHandler.Put)
	v1.DELETE("/role/:id", jwt.Verify, roleHandler.Delete)
}
