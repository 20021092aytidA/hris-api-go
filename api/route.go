package api

import (
	"go-hris/api/request"
	"go-hris/api/role"
	"go-hris/api/user"
	"go-hris/api/userdetail"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	v1 := c.Group("hris-api/v1")

	user.InitRoute(v1)
	request.InitRoute(v1)
	role.InitRoute(v1)
	userdetail.InitRoute(v1)
}
