package userdetail

import (
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1 *gin.RouterGroup) {
	var userDetailHandler Handler

	v1.GET("/user-detail", userDetailHandler.Get, jwt.Verify)
	v1.POST("/user-detail", userDetailHandler.Post, jwt.Verify)
	v1.PUT("/user-detail/:id", userDetailHandler.Put, jwt.Verify)
	v1.DELETE("/user-detail/:id", userDetailHandler.Delete, jwt.Verify)
}
