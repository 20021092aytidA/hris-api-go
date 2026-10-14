package role

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

var roleService Service

func (role Handler) Get(c *gin.Context) {
	var listRoles []View
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
			"message": "failed retrieving roles!",
			"error":   errParam.Error(),
		})
		return
	}

	listRoles, err = roleService.Find(param)
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

func (role Handler) Post(c *gin.Context) {
	var newRole Create
	if err := c.ShouldBindJSON(&newRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed creating role!",
			"error":   "missing body!",
		})
		return
	}

	newRole.CreatedAt = time.Now()
	if err := roleService.Create(newRole); err != nil {
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

func (role Handler) Put(c *gin.Context) {
	id := c.Param("id")

	var newRole Update
	if err := c.ShouldBindJSON(&newRole); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed updating role!",
			"error":   "missing body or param",
		})
		return
	}

	newRole.UpdatedAt = time.Now()
	if err := roleService.Update(id, newRole); err != nil {
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

func (role Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed deleting role!",
			"error":   "missing param",
		})
		return
	}

	if err := roleService.Erase(id); err != nil {
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
