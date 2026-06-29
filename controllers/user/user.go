package user

import (
	"errors"
	"go-hris/middleware/jwt"
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

	var param usermodel.AllParam
	param.Pagination.Page = 1
	param.Pagination.Limit = 10

	if errParam := c.ShouldBindQuery(&param); errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed retrieving users!",
			"error":   errParam.Error(),
		})
		return
	}

	listUsers, err = userservice.Find(param)
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

	var createdUser *usermodel.Create
	var errCreation error
	if createdUser, errCreation = userservice.Create(&newUser); errCreation != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed creating new user!",
			"error":   errCreation.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "user created!",
		"data":    createdUser,
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
	paramForPassword := usermodel.AllParam{
		AllowedParam: usermodel.AllowedParam{
			Username: *userLogin.Username,
		},
	}
	currUser, getCurrErr := userservice.FindForPassword(paramForPassword)
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

	//TOKEN
	jsonWebToken, jwtErr := jwt.Create(*currUser.Id, *currUser.RoleID)
	if jwtErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed logging in!",
			"error":   "jwt creation failed",
		})
		return
	}

	c.SetCookie("jwt", jsonWebToken, int(24*time.Hour), "/", "localhost", true, true)

	var param usermodel.AllParam
	param.AllowedParam.Username = *userLogin.Username

	currUserWithoutPass, errFind := userservice.Find(param)
	if errFind != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed finding user!",
			"error":   errFind.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "user logged in!",
		"currUser": currUserWithoutPass[0],
		"token":    jsonWebToken,
	})
}
