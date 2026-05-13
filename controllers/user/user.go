package user

import (
	"go-hris/helpers/request"
	usermodel "go-hris/models/user"
	userservice "go-hris/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var listUsers []usermodel.View
	var err error
	qry := c.Request.URL.RawQuery

	listUsers, err = userservice.Find(request.ProcessQry(qry))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed retrieving users!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "users retrieved!",
		"data":    listUsers,
	})
}

func Post(c *gin.Context) {
	var newUser usermodel.Create

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed creating new user!",
			"error":   "missing body",
		})
		return
	}

	if err := userservice.Create(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed creating new user!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "user created!",
		"data":    newUser,
	})
}

func Put(c *gin.Context) {
	id := c.Param("id")
	var newUser usermodel.Update

	if err := c.ShouldBindJSON(&newUser); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed updating user!",
			"error":   "missing body or param",
		})
		return
	}

	if err := userservice.Update(id, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed updating user!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user updated!",
		"data":    newUser,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed deleting user!",
			"error":   "missing param",
		})
		return
	}

	if err := userservice.Erase(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed deleting user!",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user with ID: " + id + " deleted!",
	})
}
