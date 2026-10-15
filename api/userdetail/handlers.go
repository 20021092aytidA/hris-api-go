package userdetail

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

var userDetailService Service

func (userDetail Handler) Get(c *gin.Context) {
	var listUserDetails []ViewModel
	var err error

	param := AllParam{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}

	if errParam := c.ShouldBindQuery(&param); errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed retrieving user details",
			"error":   errParam.Error(),
		})
		return
	}

	listUserDetails, err = userDetailService.Find(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed retrieving user details!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user details retrieved!",
		"data":    listUserDetails,
	})

}

func (userDetail Handler) Post(c *gin.Context) {
	var newUserDetail CreateModel

	if err := c.ShouldBindJSON(&newUserDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed creating new user detail!",
			"error":   "missing body",
		})
		return
	}

	newUserDetail.CreatedAt = time.Now()
	if err := userDetailService.Create(newUserDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed creating new user detail!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "new user detail created!",
		"data":    newUserDetail,
	})
}

func (userDetail Handler) Put(c *gin.Context) {
	id := c.Param("id")
	var newUserDetail UpdateModel

	if err := c.ShouldBindJSON(&newUserDetail); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed updating user detail!",
			"error":   "missing body or param",
		})
		return
	}

	newUserDetail.UpdatedAt = time.Now()
	if err := userDetailService.Update(id, newUserDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed updating user detail!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user detail updated!",
		"data":    newUserDetail,
	})
}

func (userDetail Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed deleting user detail!",
			"error":   "missing param",
		})
		return
	}

	if err := userDetailService.Erase(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed deleting user detail!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"messasge": "user detail with ID: " + id + " deleted!",
	})
}
