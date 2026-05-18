package jwt

import (
	"errors"
	"fmt"
	"go-hris/config/env"
	"time"

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

func Verify(tokenString string) (*jwt.Token, error) {
	jwtKey := []byte(env.ENV.JWTKey)
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaim{}, func(t *jwt.Token) (any, error) {
		// Ensure the signing method is what you expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", t.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("INVALID TOKEN")
	}

	return token, nil
}
