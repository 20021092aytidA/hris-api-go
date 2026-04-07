package jwthelper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBearerToken(header string) string {
	if header != "" {
		return header[len("Bearer "):]
	}

	return ""
}

func CheckAndValidateToken(c *gin.Context, typeReq string) bool {
	var bearerToken string = getBearerToken(c.GetHeader("Authorization"))
	if bearerToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":      http.StatusUnauthorized,
			"message":     fmt.Sprintf("Failed to update %s!", typeReq),
			"description": "No token!",
		})
		return false
	}

	if verifyToken := VerifyToken(bearerToken); verifyToken != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":      http.StatusUnauthorized,
			"message":     fmt.Sprintf("Failed to update %s!", typeReq),
			"description": "Invalid token!",
		})
		return false
	}

	return true
}
