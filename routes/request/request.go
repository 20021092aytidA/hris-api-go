package request

import (
	"go-hrs/controllers/request"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	//VERSION #1
	c.GET("hris-api/v1/request", request.Get)
	c.POST("hris-api/v1/request", request.Post)
	// c.DELETE("hris-api/v1/request/:id")
}
