package userdetail

import (
	"go-hris/controllers/userdetail"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	// VERSION #1
	c.GET("hris-api/v1/user-detail", userdetail.Get)
	c.POST("hris-api/v1/user-detail", userdetail.Post)
	c.PUT("hris-api/v1/user-detail/:id", userdetail.Put)
	c.DELETE("hris-api/v1/user-detail/:id", userdetail.Delete)
}
