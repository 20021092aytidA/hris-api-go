package usercontroller

import (
	"fmt"
	"go-hrs/helpers/jwthelper"
	"go-hrs/models/usermodel"
	"go-hrs/services/applicantdetailservice"
	"go-hrs/services/roleservice"
	"go-hrs/services/userservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "user")
	if !isValidJWT {
		return
	}

	var users []usermodel.ViewUser
	var err error = nil
	query := c.Request.URL.RawQuery

	users, err = userservice.GetUsers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to retrieve users!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully retrieve users!",
		"data":    users,
	})
}

func CreateUser(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "user")
	if !isValidJWT {
		return
	}

	var user usermodel.CreateUser
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to create new user!",
			"description": "Missing body!",
		})
		return
	}

	//ROLE
	role, _ := roleservice.GetRoles(fmt.Sprintf("role_id=%v", *user.RoleID))
	if len(role) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      http.StatusNotFound,
			"message":     "Failed to create new user!",
			"description": "Role not found!",
		})
		return
	}

	//APPLICANT DETAIL
	if user.ApplicantDetailID != nil {
		appDet, _ := applicantdetailservice.GetApplicantDetails(fmt.Sprintf("applicant_detail_id=%v", *user.ApplicantDetailID))
		if len(appDet) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":      http.StatusNotFound,
				"message":     "Failed to create new user!",
				"description": "Applciant detail not found!",
			})
			return
		}
	}

	if err := userservice.CreateUser(user); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new user!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created new user!",
	})
}

func UpdateUser(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "user")
	if !isValidJWT {
		return
	}

	id := c.Param("id")
	var user usermodel.UpdateUser

	if err := c.ShouldBind(&user); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to update user!",
			"description": "Missing body or param!",
		})
		return
	}

	//EXIST
	existUser, _ := userservice.GetUsers(fmt.Sprintf("user_id=%v", id))
	if len(existUser) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      http.StatusNotFound,
			"message":     "Failed to update user!",
			"description": "User not found!",
		})
		return
	}

	//ROLE
	if user.RoleID != nil {
		role, _ := roleservice.GetRoles(fmt.Sprintf("role_id=%v", *user.RoleID))
		if len(role) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":      http.StatusNotFound,
				"message":     "Failed to update user!",
				"description": "Role not found!",
			})
			return
		}
	}

	//APPLICANT DETAIL
	if user.ApplicantDetailID != nil {
		appDet, _ := applicantdetailservice.GetApplicantDetails(fmt.Sprintf("applicant_detail_id=%v", *user.ApplicantDetailID))
		if len(appDet) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":      http.StatusNotFound,
				"message":     "Failed to update user!",
				"description": "Applciant detail not found!",
			})
			return
		}
	}

	if err := userservice.UpdateUser(id, user); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to update user!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully updated user!",
	})

}

func DeleteUser(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "user")
	if !isValidJWT {
		return
	}

	deletedBy := c.Query("deleted_by")
	id := c.Param("id")

	if id == "" || deletedBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to delete user!",
			"description": "Missing param or query!",
		})
		return
	}

	//EXIST
	user, _ := userservice.GetUsers(fmt.Sprintf("user_id=%v", id))
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      http.StatusNotFound,
			"message":     "Failed to delete user!",
			"description": "User not found!",
		})
	}

	if err := userservice.DeleteUser(id, deletedBy); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to delete user!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully deleted user!",
	})

}
