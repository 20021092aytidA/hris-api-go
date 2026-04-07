package adminroute

import (
	"go-hrs/controllers/admincontroller"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	c.GET("/hrs-api/admins", admincontroller.GetAdmins)
	c.POST("/hrs-api/admin", admincontroller.CreateAdmin)
	c.POST("/hrs-api/admin/login", admincontroller.LoginAdmin)
	c.PUT("/hrs-api/admin/:id", admincontroller.UpdateAdmin)
	c.DELETE("/hrs-api/admin/:id", admincontroller.DeleteAdmin)
}
