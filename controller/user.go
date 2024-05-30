package controller

import (
	"fmt"
	"net/http"

	"github.com/ZEQUANR/zhulong/logger"
	"github.com/ZEQUANR/zhulong/model"
	"github.com/ZEQUANR/zhulong/model/api"
	"github.com/ZEQUANR/zhulong/services"
	"github.com/ZEQUANR/zhulong/utils"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	data := &api.Login{}
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

	c.JSON(http.StatusOK, &gin.H{"token": token})
}

func UserInfo(c *gin.Context) {
	userId, err := utils.ParseUserIDInToken(c)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
		return
	}

	user, err := services.QueryUserById(userId)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
		return
	}

	if user.Role == model.Admin {
		info, err := services.QueryAdministratorsById(userId)
		if err != nil {
			logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"userId":  user.ID,
			"account": user.Account,
			"avatar":  info.Avatar,
			"role":    user.Role,
			"name":    info.Name,
			"college": info.College,
			"phone":   info.Phone,
			"number":  info.Number,
		})

		return
	}

	if user.Role == model.Teacher {
		info, err := services.QueryTeachersById(userId)
		if err != nil {
			logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"userId":  user.ID,
			"account": user.Account,
			"avatar":  info.Avatar,
			"role":    user.Role,
			"name":    info.Name,
			"college": info.College,
			"phone":   info.Phone,
			"number":  info.Number,
		})

		return
	}

	if user.Role == model.Student {
		info, err := services.QueryStudentsById(userId)
		if err != nil {
			logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"userId":  user.ID,
			"account": user.Account,
			"avatar":  info.Avatar,
			"role":    user.Role,
			"name":    info.Name,
			"college": info.College,
			"phone":   info.Phone,
			"number":  info.Number,
		})

		return
	}

	logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
}

func UserTeacherList(c *gin.Context) {
	userId, err := utils.ParseUserIDInToken(c)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
		return
	}

	user, err := services.QueryUserById(userId)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
		return
	}

	if user.Role == model.Student {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionQuery, logger.ErrorBodyPermissions, fmt.Errorf(""))
		return
	}

	result, err := services.QueryTeacherList()
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyTeacherList, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func UserEditor(c *gin.Context) {
	// userId, err := utils.ParseUserIDInToken(c)
	// if err != nil {
	// 	logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
	// 	return
	// }

	// user, err := services.QueryUserById(userId)
	// if err != nil {
	// 	logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
	// 	return
	// }

	// if user.Role == model.Admin {
	// 	data := api.Administrator{}

	// 	if err := c.BindJSON(&data); err != nil {
	// 		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
	// 		return
	// 	}

	// 	result, err := services.UpdateAdministratorsById(user.ID, data)
	// 	if err != nil {
	// 		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"user_id":  user.ID,
	// 		"account":  user.Account,
	// 		"role":     user.Role,
	// 		"name":     result.Name,
	// 		"identity": result.Identity,
	// 		"college":  result.College,
	// 		"phone":    result.Phone,
	// 	})

	// 	return
	// }

	// if user.Role == model.Teacher {
	// 	data := api.Teacher{}

	// 	if err := c.BindJSON(&data); err != nil {
	// 		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
	// 		return
	// 	}

	// 	result, err := services.UpdateTeachersById(user.ID, data)
	// 	if err != nil {
	// 		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"user_id":  user.ID,
	// 		"account":  user.Account,
	// 		"role":     user.Role,
	// 		"name":     result.Name,
	// 		"identity": result.Identity,
	// 		"college":  result.College,
	// 		"phone":    result.Phone,
	// 	})

	// 	return
	// }

	// if user.Role == model.Student {
	// 	data := api.Student{}

	// 	if err := c.BindJSON(&data); err != nil {
	// 		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
	// 		return
	// 	}

	// 	result, err := services.UpdateStudentsById(user.ID, data)
	// 	if err != nil {
	// 		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"user_id":  user.ID,
	// 		"account":  user.Account,
	// 		"role":     user.Role,
	// 		"name":     result.Name,
	// 		"college":  result.College,
	// 		"phone":    result.Phone,
	// 		"subject":  result.Subject,
	// 		"class":    result.Class,
	// 		"identity": result.Identity,
	// 	})

	// 	return
	// }

	// logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyQueryingUser, err)
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
