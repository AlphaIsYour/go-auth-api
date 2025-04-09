package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecretkey")

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}

	username := claims["username"].(string)
	return username, nil
}
