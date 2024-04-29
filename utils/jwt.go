package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Load key from somewhere, for example an environment variable
const secret = "Jerry"

func GenerateToken(userID uint) (string, error) {
	const tokenExpireDuration = time.Hour * 24
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(tokenExpireDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParseAToken(tokenString string, field string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims[field] != nil {
			return claims[field], nil
		} else {
			return nil, fmt.Errorf("unexpected signature field: %s", field)
		}
	} else {
		return nil, fmt.Errorf("unexpected signature field: %s", field)
	}
}

func ExtractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", fmt.Errorf("the Authorization field cannot be found from the request header")
}
