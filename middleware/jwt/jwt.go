package jwt

import (
	"fmt"
	"go-hris/config/env"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaim struct {
	jwt.RegisteredClaims
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
}

func Create(userID int, roleID int) (string, error) {
	jwtKey := []byte(env.ENV.JWTKey)
	duration := 24 * time.Hour
	method := jwt.SigningMethodHS256

	claim := jwtClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "hris-api-go",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		UserID: userID,
		RoleID: roleID,
	}

	token := jwt.NewWithClaims(method, claim)
	return token.SignedString(jwtKey)
}

func Verify(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "request failed!",
			"error":   "token required",
		})

		c.Abort()
		return
	}

	jwtTokenArr := strings.Split(bearer, " ")

	if len(jwtTokenArr) < 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "request failed!",
			"error":   "token required",
		})

		c.Abort()
		return
	}

	jwtToken := jwtTokenArr[1]

	jwtKey := []byte(env.ENV.JWTKey)
	token, err := jwt.ParseWithClaims(jwtToken, &jwtClaim{}, func(t *jwt.Token) (any, error) {
		// Ensure the signing method is what you expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method (%v)", t.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "failed verification jwt!",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": "failed verification jwt!",
			"error":   "invalid token",
		})
		c.Abort()
		return
	}

	url := c.Request.URL.String()
	splitURL := strings.Split(url, "/")
	if splitURL[len(splitURL)-1] == "token" {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "verification jwt successful!",
		})
		c.Next()
	} else {
		c.Next()
	}
}
