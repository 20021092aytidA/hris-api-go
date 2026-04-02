package rolecontroller

import (
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
