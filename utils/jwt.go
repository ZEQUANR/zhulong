package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Load key from somewhere, for example an environment variable
const (
	secret = "Jerry"
	UserId = "user_id"
)

func GenerateToken(userID uint) (string, error) {
	const tokenExpireDuration = time.Hour * 24
	claims := jwt.MapClaims{
		"authorized": true,
		UserId:       userID,
		"exp":        time.Now().Add(tokenExpireDuration).Unix(),
	}

	str := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := str.SignedString([]byte(secret))
	if err != nil {
		return token, fmt.Errorf("func: GenerateToken | index: 0 | err: %w", err)
	}

	return token, nil
}

func ParseAToken(tokenString string, field string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("func: ParseAToken | index: 0 | err: unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("func: ParseAToken | index: 1 | err: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims[field] != nil {
			return claims[field], nil
		}
	}

	return nil, fmt.Errorf("func: ParseAToken | index: 2 | err: unexpected signature field: %s", field)
}

func ExtractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", fmt.Errorf("func: ExtractToken | index: 0 | err: the Authorization field cannot be found from the request header")
}

func ParseUserIDInToken(c *gin.Context) (int, error) {
	token, err := ExtractToken(c)
	if err != nil {
		return 0, err
	}

	id, err := ParseAToken(token, UserId)
	if err != nil {
		return 0, err
	}

	return int(id.(float64)), nil
}
