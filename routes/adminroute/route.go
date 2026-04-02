package adminroute

import (
	"go-hrs/controllers/admincontroller"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	c.GET("/hrs-api/admins", admincontroller.GetAdmins)
}
