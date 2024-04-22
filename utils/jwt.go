package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Load key from somewhere, for example an environment variable
const secret = "Jerry"

func GenerateToken(userID uint, unit int) (string, error) {
	const tokenExpireDuration = time.Hour * 24
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"unit":       unit,
		"exp":        time.Now().Add(tokenExpireDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParseAToken(tokenString string, field string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims[field], nil
	} else {
		return nil, fmt.Errorf("Unexpected signature field: %s", field)
	}
}
