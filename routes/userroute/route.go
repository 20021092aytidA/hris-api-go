package userroute

import (
	"go-hrs/controllers/usercontroller"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	c.GET("/hrs-api/users", usercontroller.GetUsers)
}
