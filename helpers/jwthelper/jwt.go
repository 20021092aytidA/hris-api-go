package jwthelper

import (
	"errors"
	"fmt"
	"go-hrs/config/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaim struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func CreateJWTToken(username string) (string, error) {
	jwtKey := []byte(env.ENV.JWTKey)
	duration := 24 * time.Hour
	method := jwt.SigningMethodHS256

	claim := jwtClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "pacific",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(method, claim)
	return token.SignedString(jwtKey)

}

// verifyToken parses and validates the JWT string
func VerifyToken(tokenString string) error {
	jwtKey := []byte(env.ENV.JWTKey)
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaim{}, func(t *jwt.Token) (any, error) {
		// Ensure the signing method is what you expect
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", t.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("INVALID TOKEN")
	}

	return nil
}
