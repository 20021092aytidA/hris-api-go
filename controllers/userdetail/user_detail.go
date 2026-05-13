package userdetail

import (
	"go-hris/helpers/request"
	userdetailmodel "go-hris/models/userdetail"
	userdetailservice "go-hris/services/userdetail"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var listUserDetails []userdetailmodel.View
	var err error
	qry := c.Request.URL.RawQuery

	listUserDetails, err = userdetailservice.Find(request.ProcessQry(qry))
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

func Post(c *gin.Context) {
	var newUserDetail userdetailmodel.Create

	if err := c.ShouldBindJSON(&newUserDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed creating new user detail!",
			"error":   "missing body",
		})
		return
	}

	newUserDetail.CreatedAt = time.Now()
	if err := userdetailservice.Create(newUserDetail); err != nil {
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

func Put(c *gin.Context) {
	id := c.Param("id")
	var newUserDetail userdetailmodel.Update

	if err := c.ShouldBindJSON(&newUserDetail); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed updating user detail!",
			"error":   "missing body or param",
		})
		return
	}

	newUserDetail.UpdatedAt = time.Now()
	if err := userdetailservice.Update(id, newUserDetail); err != nil {
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

func Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed deleting user detail!",
			"error":   "missing param",
		})
		return
	}

	if err := userdetailservice.Erase(id); err != nil {
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
