package usercontroller

import (
	"fmt"
	"go-hrs/models/usermodel"
	"go-hrs/services/roleservice"
	"go-hrs/services/userservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []usermodel.ViewUser
	var err error = nil
	query := c.Request.URL.RawQuery

	users, err = userservice.GetUsers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to retrieve users!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully retrieve users!",
		"data":    users,
	})
}

func CreateUser(c *gin.Context) {
	var user usermodel.CreateUser
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to create new user!",
			"description": "Missing body!",
		})
		return
	}

	//ROLE
	role, _ := roleservice.GetRoles(fmt.Sprintf("role_id=%v", *user.RoleID))
	if len(role) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      http.StatusNotFound,
			"message":     "Failed to create new user!",
			"description": "Role not found!",
		})
		return
	}

	if err := userservice.CreateUser(user); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new user!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created new user!",
	})
}
