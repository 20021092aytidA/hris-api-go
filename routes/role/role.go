package role

import (
	"go-hris/controllers/role"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	// VERSION #1
	c.GET("hris-api/v1/role", role.Get)
	c.POST("hris-api/v1/role", role.Post)
	c.PUT("hris-api/v1/role/:id", role.Put)
	c.DELETE("hris-api/v1/role/:id", role.Delete)
}
