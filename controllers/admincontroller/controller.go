package admincontroller

import (
	"fmt"
	"go-hrs/helpers/jwthelper"
	"go-hrs/models/adminmodel"
	"go-hrs/services/adminservice"
	"go-hrs/services/roleservice"
	"go-hrs/services/userservice"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAdmins(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "role")
	if !isValidJWT {
		return
	}

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

func CreateAdmin(c *gin.Context) {
	var create adminmodel.CreateAdmin
	if err := c.ShouldBind(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to create new admin!",
			"description": "Missing body!",
		})
		return
	}

	//CHECK AVAILABLE USER
	user, _ := userservice.GetUsers(fmt.Sprintf("user_id=%v", *create.UserID))
	if len(user) == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new admin!",
			"description": "User not found!",
		})
		return
	}

	//CHECK AVAILABLE ROLE
	role, _ := roleservice.GetRoles(fmt.Sprintf("role_id=%v", *create.RoleID))
	if len(role) == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new admin!",
			"description": "Role not found!",
		})
		return
	}

	//DUPLICATE USERNAME
	admins, _ := adminservice.GetAdmins(fmt.Sprintf("username=%s", *create.Username))
	if len(admins) != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new admin!",
			"description": "Duplicate username!",
		})
		return
	}

	encryptPass, errBcrypt := bcrypt.GenerateFromPassword([]byte(*create.Password), bcrypt.DefaultCost)
	if errBcrypt != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to encrpyt password!",
			"description": errBcrypt.Error(),
		})
		return
	}
	stringEncryptPass := string(encryptPass)
	create.Password = &stringEncryptPass

	if err := adminservice.CreateAdmin(create); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new admin!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created new admin!",
	})
}

func UpdateAdmin(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "admin")
	if !isValidJWT {
		return
	}

	adminID := c.Param("id")
	var update adminmodel.UpdateAdmin

	if err := c.ShouldBind(&update); err != nil || adminID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to update admin!",
			"description": "Missing body or param!",
		})
		return
	}

	//CHECK IF EXIST
	admin, _ := adminservice.GetAdmins(fmt.Sprintf("admin_id=%s", adminID))
	if len(admin) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      http.StatusNotFound,
			"message":     "Failed to update admin!",
			"description": "Admin not found!",
		})
		return
	}

	//CHECK AVAILABLE USER
	if update.UserID != nil {
		user, _ := userservice.GetUsers(fmt.Sprintf("user_id=%v", *update.UserID))
		if len(user) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":      http.StatusNotFound,
				"message":     "Failed to update admin!",
				"description": "User not found!",
			})
			return
		}
	}

	//CHECK AVAILABLE ROLE
	if update.RoleID != nil {
		role, _ := roleservice.GetRoles(fmt.Sprintf("role_id=%v", *update.RoleID))
		if len(role) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":      http.StatusNotFound,
				"message":     "Failed to update admin!",
				"description": "Role not found!",
			})
			return
		}
	}

	//DUPLICATE USERNAME
	if update.Username != nil {
		admins, _ := adminservice.GetAdmins(fmt.Sprintf("username=%s", *update.Username))
		if len(admins) != 0 {
			c.JSON(http.StatusConflict, gin.H{
				"status":      http.StatusConflict,
				"message":     "Failed to update admin!",
				"description": "Duplicate username!",
			})
			return
		}
	}

	if update.Password != nil {
		admins, _ := adminservice.GetAdmins(fmt.Sprintf("username=%s", *update.Username))
		if err := bcrypt.CompareHashAndPassword([]byte(*admins[0].Password), []byte(*update.Password)); err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"status":      http.StatusConflict,
				"message":     "Failed to update admin!",
				"description": "Wrong password!",
			})
			return
		}

		encryptedPass, bcryptErr := bcrypt.GenerateFromPassword([]byte(*update.Password), bcrypt.DefaultCost)
		if bcryptErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":      http.StatusInternalServerError,
				"message":     "Failed to encrypt password!",
				"description": bcryptErr.Error(),
			})
			return
		}
		var encryptPassString string = string(encryptedPass)
		update.Password = &encryptPassString
	}

	if errUpdate := adminservice.UpdateAdmin(adminID, update); errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to update admin!",
			"description": errUpdate.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully updated admin!",
	})
}

func DeleteAdmin(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "admin")
	if !isValidJWT {
		return
	}

	adminID := c.Param("id")
	deletedBy := c.Query("deleted_by")
	if adminID == "" || deletedBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to delete admin!",
			"description": "Missing param or query!",
		})
		return
	}

	admin, _ := adminservice.GetAdmins(fmt.Sprintf("admin_id=%s", adminID))
	if len(admin) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      http.StatusNotFound,
			"message":     "Failed to delete admin!",
			"description": "No admin found!",
		})
		return
	}

	if errDel := adminservice.DeleteAdmin(adminID, deletedBy); errDel != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to delete admin!",
			"description": errDel.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully deleted admin!",
	})
}

func LoginAdmin(c *gin.Context) {
	var login adminmodel.LoginAdmin

	//NO BODY
	if err := c.ShouldBind(&login); err != nil {
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
