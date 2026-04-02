package admincontroller

import (
	"go-hrs/models/adminmodel"
	"go-hrs/services/adminservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdmins(c *gin.Context) {
	var admins []adminmodel.ViewAdmin
	var err error = nil
	query := c.Request.URL.RawQuery

	admins, err = adminservice.GetAdmins(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to retrieve admins!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully retrieve admins!",
		"data":    admins,
	})
}
