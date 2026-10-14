package user

import (
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1 *gin.RouterGroup) {
	var userHandler Handler

	//NEED JWT
	v1.GET("/user", userHandler.Get, jwt.Verify)
	v1.PUT("/user/:id", userHandler.Put, jwt.Verify)
	v1.DELETE("/user/:id", userHandler.Delete, jwt.Verify)
	v1.POST("/user/token", jwt.Verify)

	//NO NEED JWT
	v1.POST("/user", userHandler.Post)
	v1.POST("/user/login", userHandler.Login)
}
