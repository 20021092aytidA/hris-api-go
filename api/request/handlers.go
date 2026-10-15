package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

var requestService Service

func (req Handler) Get(c *gin.Context) {
	var listRequest []ViewModel
	var err error

	var param AllParam
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

	listRequest, err = requestService.Find(param)

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

func (req Handler) Post(c *gin.Context) {
	var newRequest CreateModel
	if err := c.ShouldBindJSON(&newRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed creating new request!",
			"error":   "missing body!",
		})
		return
	}

	if err := requestService.Create(newRequest); err != nil {
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

func (req Handler) Put(c *gin.Context) {
	id := c.Param("id")

	var newRequest UpdateModel
	if err := c.ShouldBindJSON(&newRequest); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed updating role!",
			"error":   "missing body or param",
		})
		return
	}

	if err := requestService.Update(id, newRequest); err != nil {
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

func (req Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed deleting request!",
			"error":   "missing param",
		})
		return
	}

	if err := requestService.Erase(id); err != nil {
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
