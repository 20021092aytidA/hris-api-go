package user

import (
	"errors"
	"fmt"
	"go-hris/helpers/request"
	usermodel "go-hris/models/user"
	userservice "go-hris/services/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	newUser.CreatedAt = time.Now()
	//PASSWORD
	bytePass := []byte(*newUser.Password)
	byteHashedPass, errHash := bcrypt.GenerateFromPassword(bytePass, 10)
	if errHash != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed creating new user!",
			"error":   "password hashing\n" + errHash.Error(),
		})
		return
	}
	*newUser.Password = string(byteHashedPass)

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

	newUser.UpdatedAt = time.Now()
	//PASSWORD
	if newUser.Password != nil {
		var customQry = make(map[string]any)
		customQry["id"] = id

		currUser, getUserErr := userservice.FindForPassword(request.ProcessQry(fmt.Sprintf("id=%s", id)))
		if getUserErr != nil {
			if errors.Is(getUserErr, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": "failed updating user!",
					"error":   "user not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "failed updating user!",
				"error":   "password validation failed (GET)",
			})
			return
		}

		byteCurrPass := []byte(*currUser.Password)
		byteConfirmPass := []byte(*newUser.ConfirmPassword)
		byteNewPass := []byte(*newUser.Password)

		errHash := bcrypt.CompareHashAndPassword(byteCurrPass, byteConfirmPass)
		if errHash != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "failed to update user!",
				"error":   "incorrect password",
			})
			return
		}

		hashNewPass, hashErr := bcrypt.GenerateFromPassword(byteNewPass, 10)
		if hashErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "failed to update user!",
				"error":   "hashing password",
			})
			return
		}

		*newUser.Password = string(hashNewPass)
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

func Login(c *gin.Context) {
	var userLogin usermodel.Login

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "falied logging in!",
			"error":   "missign body",
		})
		return
	}

	byteLoginPass := []byte(*userLogin.Password)
	currUser, getCurrErr := userservice.FindForPassword(request.ProcessQry(fmt.Sprintf("username=%s", *userLogin.Username)))
	if getCurrErr != nil {
		if errors.Is(getCurrErr, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "failed logging in!",
				"error":   "user not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed logging in!",
			"error":   "failed validating password (GET)",
		})
		return
	}

	byteHashLoginPass := []byte(*currUser.Password)
	validErr := bcrypt.CompareHashAndPassword(byteHashLoginPass, byteLoginPass)
	if validErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "failed logging in!",
			"error":   "incorrect password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user logged in!",
		"token":   "",
	})
}
