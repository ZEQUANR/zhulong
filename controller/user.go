package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ZEQUANR/zhulong/logger"
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
		logger.CreateLog(c, logger.ErrorGroupWarning, logger.ErrorWhoClient, logger.ErrorTypeParameter, logger.ErrorBodyDataFormat, err)
		return
	}

	user, err := services.QueryUserAccountPassword(data.Account, utils.Md5Encode(data.Password))
	if err != nil {
		logger.CreateLog(c, logger.ErrorGroupWarning, logger.ErrorWhoClient, logger.ErrorTypeParameter, logger.ErrorBodyDatabase, err)
		return
	}

	token, err := utils.GenerateToken(uint(user.ID))
	if err != nil {
		logger.CreateLog(c, logger.ErrorGroupWarning, logger.ErrorWhoClient, logger.ErrorTypeParameter, logger.ErrorBodyDataFormat, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}

func UserInfo(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) != 2 {
		logger.CreateLog(c, logger.ErrorGroupWarning, logger.ErrorWhoClient, logger.ErrorTypeParameter, logger.ErrorBodyDataFormat, fmt.Errorf("Failed querying user"))
		return
	}

	token := strings.Split(bearerToken, " ")[1]
	id, err := utils.ParseAToken(token, "user_id")
	if err != nil {
		logger.CreateLog(c, logger.ErrorGroupWarning, logger.ErrorWhoClient, logger.ErrorTypeParameter, logger.ErrorBodyDataFormat, err)
		return
	}

	user, err := services.QueryUserById(int(id.(float64)))
	if err != nil {
		logger.CreateLog(c, logger.ErrorGroupWarning, logger.ErrorWhoClient, logger.ErrorTypeParameter, logger.ErrorBodyDataFormat, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user_id": user.ID,
			"account": user.Account,
			"role":    user.Role,
		},
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
