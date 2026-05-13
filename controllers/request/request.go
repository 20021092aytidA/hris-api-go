package request

import (
	requesthelper "go-hris/helpers/request"
	"go-hris/models/request"
	requestservice "go-hris/services/request"
	"net/http"
	"time"

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

func Post(c *gin.Context) {
	var newRequest request.Create
	if err := c.ShouldBindJSON(&newRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed creating new request!",
			"error":   "missing body!",
		})
		return
	}

	newRequest.CreatedAt = time.Now()
	if err := requestservice.Create(newRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed creating new request!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "new request created!",
		"data":    newRequest,
	})

}

func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed deleting request!",
			"error":   "missing param",
		})
		return
	}

	if err := requestservice.Erase(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed deleting request!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success request deletion!",
	})
}
