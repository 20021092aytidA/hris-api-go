package rolecontroller

import (
	"fmt"
	"go-hrs/models/rolemodel"
	"go-hrs/services/roleservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	var roles []rolemodel.ViewRole
	var err error = nil
	query := c.Request.URL.RawQuery

	roles, err = roleservice.GetRoles(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to retrieve roles!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully retrieve roles!",
		"data":    roles,
	})
}

func CreateRole(c *gin.Context) {
	var role rolemodel.CreateRole
	var checkRole []rolemodel.ViewRole

	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to create new role!",
			"description": "Missing body parameter!",
		})
		return
	}

	checkRole, _ = roleservice.GetRoles(fmt.Sprintf("role_name=%s", *role.RoleName))
	if len(checkRole) != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new role!",
			"description": "Duplicate role name!",
		})
		return
	}

	if err := roleservice.CreateRole(role); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new role!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created new role!",
	})
}
