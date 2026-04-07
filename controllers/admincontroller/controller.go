package admincontroller

import (
	"fmt"
	"go-hrs/helpers/jwthelper"
	"go-hrs/models/adminmodel"
	"go-hrs/services/adminservice"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAdmins(c *gin.Context) {
	var admins []adminmodel.ViewAdmin
	var err error = nil
	query := c.Request.URL.RawQuery

	admins, err = adminservice.GetAdmins(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to retrieve admins!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully retrieve admins!",
		"data":    admins,
	})
}

func LoginAdmin(c *gin.Context) {
	var login adminmodel.LoginAdmin

	//NO BODY
	if err := c.ShouldBind(&login).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to login!",
			"description": "Missing body!",
		})
		return
	}

	//NOT FOUND
	admin, _ := adminservice.GetAdmins(fmt.Sprintf("username=%s", *login.Username))
	if len(admin) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Failed to login!", "description": "No user found!"})
		return
	}

	//WRONG PASSWORD
	passwordErr := bcrypt.CompareHashAndPassword([]byte(*admin[0].Password), []byte(*login.Password))
	if passwordErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":      http.StatusUnauthorized,
			"message":     "Failed to login!",
			"description": "Wrong password!",
		})
		return
	}

	//JWT
	jwtToken, errJWT := jwthelper.CreateJWTToken(*login.Username)
	if errJWT != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to create jwt token!",
			"description": errJWT.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successful login!",
		"token":   jwtToken,
	})
}
