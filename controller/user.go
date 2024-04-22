package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var data struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameter",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": "2333",
	})
}

func UserRegister(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserLogout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
