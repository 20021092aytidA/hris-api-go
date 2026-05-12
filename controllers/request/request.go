package request

import (
	requesthelper "go-hrs/helpers/request"
	"go-hrs/models/request"
	requestservice "go-hrs/services/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var listRequest []request.View
	var err error

	qry := requesthelper.ProcessQry(c.Request.URL.RawQuery)
	listRequest, err = requestservice.Find(qry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed retrieving request!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "retrieved request!",
		"data":    listRequest,
	})
}
