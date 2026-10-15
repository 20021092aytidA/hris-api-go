package user

import (
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1 *gin.RouterGroup) {
	var userHandler Handler

	//NO NEED JWT
	v1.POST("/user", userHandler.Post)
	v1.POST("/user/login", userHandler.Login)

	//NEED JWT
	v1.GET("/user", jwt.Verify, userHandler.Get)
	v1.PUT("/user/:id", jwt.Verify, userHandler.Put)
	v1.DELETE("/user/:id", jwt.Verify, userHandler.Delete)
	v1.POST("/user/token", jwt.Verify)

}
