package request

import (
	"go-hris/models/request"
	requestservice "go-hris/services/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var listRequest []request.View
	var err error

	var param request.AllParam
	param.Pagination.Page = 1
	param.Pagination.Limit = 10

	if errParam := c.ShouldBindQuery(&param); errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed retrieving request!",
			"error":   errParam.Error(),
		})
		return
	}

	listRequest, err = requestservice.Find(param)

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

func Put(c *gin.Context) {
	id := c.Param("id")

	var newRequest request.Update
	if err := c.ShouldBindJSON(&newRequest); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed updating role!",
			"error":   "missing body or param",
		})
		return
	}

	if err := requestservice.Update(id, newRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed updating request!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "request updated!",
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
		"message": "request with ID: " + id + " deleted!",
	})
}
