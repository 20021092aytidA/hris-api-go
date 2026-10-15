package userdetail

import (
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1 *gin.RouterGroup) {
	var userDetailHandler Handler

	v1.GET("/user-detail", jwt.Verify, userDetailHandler.Get)
	v1.POST("/user-detail", jwt.Verify, userDetailHandler.Post)
	v1.PUT("/user-detail/:id", jwt.Verify, userDetailHandler.Put)
	v1.DELETE("/user-detail/:id", jwt.Verify, userDetailHandler.Delete)
}
