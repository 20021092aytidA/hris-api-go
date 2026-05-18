package userdetail

import (
	"go-hris/controllers/userdetail"
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	// VERSION #1
	{
		//NEED JWT
		v1 := c.Group("hris-api/v1", jwt.Verify)
		v1.GET("/user-detail", userdetail.Get)
		v1.POST("/user-detail", userdetail.Post)
		v1.PUT("/user-detail/:id", userdetail.Put)
		v1.DELETE("/user-detail/:id", userdetail.Delete)
	}
}
