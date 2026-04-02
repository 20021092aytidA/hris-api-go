package roleroute

import (
	"go-hrs/controllers/rolecontroller"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	c.GET("/hrs-api/roles", rolecontroller.GetRoles)
	c.POST("/hrs-api/role", rolecontroller.CreateRole)
	c.PUT("/hrs-api/role/:id", rolecontroller.UpdateRole)
	c.DELETE("/hrs-api/role/:id", rolecontroller.DeleteRole)
}
