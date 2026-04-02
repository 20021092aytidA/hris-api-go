package usercontroller

import (
	"go-hrs/models/usermodel"
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
