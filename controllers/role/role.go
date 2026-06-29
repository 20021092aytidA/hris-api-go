package role

import (
	rolemodel "go-hris/models/role"
	roleservice "go-hris/services/role"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var listRoles []rolemodel.View
	var err error

	param := rolemodel.AllParam{
		Pagination: rolemodel.Pagination{
			Page:  1,
			Limit: 10,
		},
	}

	if errParam := c.ShouldBindQuery(&param); errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed retrieving roles!",
			"error":   errParam.Error(),
		})
		return
	}

	listRoles, err = roleservice.Find(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed retrieving roles!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "retrieved roles!",
		"data":    listRoles,
	})
}

func Post(c *gin.Context) {
	var newRole rolemodel.Create
	if err := c.ShouldBindJSON(&newRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed creating role!",
			"error":   "missing body!",
		})
		return
	}

	newRole.CreatedAt = time.Now()
	if err := roleservice.Create(newRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed creating new role!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "new role created!",
		"data":    newRole,
	})
}

func Put(c *gin.Context) {
	id := c.Param("id")

	var newRole rolemodel.Update
	if err := c.ShouldBindJSON(&newRole); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed updating role!",
			"error":   "missing body or param",
		})
		return
	}

	newRole.UpdatedAt = time.Now()
	if err := roleservice.Update(id, newRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed updating role!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "role updated!",
		"data":    newRole,
	})

}

func Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed deleting role!",
			"error":   "missing param",
		})
		return
	}

	if err := roleservice.Erase(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed deleting role!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "role with ID: " + id + " deleted!",
	})
}
