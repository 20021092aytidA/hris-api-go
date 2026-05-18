package role

import (
	"go-hris/controllers/role"
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	// VERSION #1
	{
		//NEED JWT
		v1 := c.Group("hris-api/v1", jwt.Verify)
		v1.GET("/role", role.Get)
		v1.POST("/role", role.Post)
		v1.PUT("/role/:id", role.Put)
		v1.DELETE("/role/:id", role.Delete)
	}
}
