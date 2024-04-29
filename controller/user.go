package controller

import (
	"net/http"

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
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
		return
	}

	user, err := services.QueryUserByAccountPassword(data.Account, utils.Md5Encode(data.Password))
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
		return
	}

	token, err := utils.GenerateToken(uint(user.ID))
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionCreate, logger.ErrorBodyCreateToken, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserInfo(c *gin.Context) {
	token, err := utils.ExtractToken(c)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyRequestHeader, err)
		return
	}

	id, err := utils.ParseAToken(token, "user_id")
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
		return
	}

	user, err := services.QueryUserById(int(id.(float64)))
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
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
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
