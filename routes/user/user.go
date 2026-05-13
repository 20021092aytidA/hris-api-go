package user

import (
	"go-hris/controllers/user"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	// VERSION #1
	c.GET("hris-api/v1/user", user.Get)
	c.POST("hris-api/v1/user", user.Post)
	c.PUT("hris-api/v1/user/:id", user.Put)
	c.DELETE("hris-api/v1/user/:id", user.Delete)
	c.POST("hris-api/v1/user/login", user.Login)
}
