package user

import (
	"go-hris/controllers/user"
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	// VERSION #1
	{
		//NEED JWT
		v1wJWT := c.Group("hris-api/v1", jwt.Verify)
		v1wJWT.GET("/user", user.Get)
		v1wJWT.PUT("/user/:id", user.Put)
		v1wJWT.DELETE("/user/:id", user.Delete)
		v1wJWT.POST("/user/token")

		//NO NEED JWT
		v1withoutJWT := c.Group("hris-api/v1")
		v1withoutJWT.POST("/user", user.Post)
		v1withoutJWT.POST("/user/login", user.Login)

	}
}
