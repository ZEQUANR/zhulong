package controller

import (
	"log/slog"
	"net/http"

	"github.com/ZEQUANR/zhulong/services"
	"github.com/ZEQUANR/zhulong/utils"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var data struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&data); err != nil {
		slog.Error(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	user, userErr := services.QueryUserById(data.Account, utils.Md5Encode(data.Password))
	if userErr != nil {
		slog.Error(userErr.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": userErr.Error(),
		})

		return
	}

	token, tokenErr := utils.GenerateToken(uint(user.ID))
	if tokenErr != nil {
		slog.Error(tokenErr.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": tokenErr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}

func UserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserRegister(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserLogout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
