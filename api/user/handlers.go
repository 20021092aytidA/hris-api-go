package user

import (
	"errors"
	"go-hris/middleware/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct{}

var userService Service

func (user Handler) Get(c *gin.Context) {
	var listUsers []ViewModel
	var err error

	var param AllParam
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

	listUsers, err = userService.Find(param)
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

func (user Handler) Post(c *gin.Context) {
	var newUser CreateModel

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed creating new user!",
			"error":   "missing body",
		})
		return
	}

	var createdUser *CreateModel
	var errCreation error
	if createdUser, errCreation = userService.Create(&newUser); errCreation != nil {
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

func (user Handler) Put(c *gin.Context) {
	id := c.Param("id")
	var newUser UpdateModel

	if err := c.ShouldBindJSON(&newUser); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed updating user!",
			"error":   "missing body or param",
		})
		return
	}

	if err := userService.Update(id, newUser); err != nil {
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

func (user Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed deleting user!",
			"error":   "missing param",
		})
		return
	}

	if err := userService.Erase(id); err != nil {
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

func (user Handler) Login(c *gin.Context) {
	var userLogin LoginModel

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "falied logging in!",
			"error":   "missign body",
		})
		return
	}

	byteLoginPass := []byte(*userLogin.Password)
	paramForPassword := AllParam{
		AllowedParam: AllowedParam{
			Username: *userLogin.Username,
		},
	}
	currUser, getCurrErr := userService.FindForPassword(paramForPassword)
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

	param := AllParam{
		AllowedParam: AllowedParam{
			Username: *userLogin.Username,
		},
		Pagination: Pagination{
			Page:  1,
			Limit: 1,
		},
	}

	currUserWithoutPass, errFind := userService.Find(param)
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
