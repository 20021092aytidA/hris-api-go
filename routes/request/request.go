package request

import (
	"go-hris/controllers/request"
	"go-hris/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	//VERSION #1
	{
		//NEED JWT
		v1 := c.Group("hris-api/v1", jwt.Verify)
		v1.GET("/request", request.Get)
		v1.POST("/request", request.Post)
		v1.DELETE("/request/:id", request.Delete)
	}
}
