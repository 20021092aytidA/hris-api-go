package rolecontroller

import (
	"fmt"
	"go-hrs/helpers/jwthelper"
	"go-hrs/models/rolemodel"
	"go-hrs/services/roleservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "role")
	if !isValidJWT {
		return
	}

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
	isValidJWT := jwthelper.CheckAndValidateToken(c, "role")
	if !isValidJWT {
		return
	}

	var role rolemodel.CreateRole
	var checkRole []rolemodel.ViewRole

	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to create new role!",
			"description": "Missing body!",
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

func UpdateRole(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "role")
	if !isValidJWT {
		return
	}

	id := c.Param("id")
	var role rolemodel.UpdateRole
	var checkRole []rolemodel.ViewRole

	if err := c.ShouldBind(&role); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to update role!",
			"description": "Missing body or param!",
		})
		return
	}

	if role.RoleName == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to update role!",
			"description": "Nothing to update!",
		})
		return
	}

	checkRole, _ = roleservice.GetRoles(fmt.Sprintf("role_name=%s", *role.RoleName))
	if len(checkRole) != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to update role!",
			"description": "Duplicate role name!",
		})
		return
	}

	if err := roleservice.UpdateRole(id, role); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to update role!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully updated role!",
	})
}

func DeleteRole(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "role")
	if !isValidJWT {
		return
	}

	id := c.Param("id")
	deletedBy := c.Query("deleted_by")

	if id == "" || deletedBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to delete role!",
			"description": "Missing param or query!",
		})
		return
	}

	if err := roleservice.DeleteRole(id, deletedBy); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to delete role!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully deleted role!",
	})
}
