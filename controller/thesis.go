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

func ThesisCreate(c *gin.Context) {
	userId, err := utils.ParseUserIDInToken(c)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
		return
	}

	data := api.CreateThesis{}
	if err := c.BindJSON(&data); err != nil {
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
		return
	}

	thesis, err := services.CreateThesis(userId, data)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionCreate, logger.ErrorBodyCreateThesis, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"thesisId": thesis.ID,
	})
}

func ThesisUpload(c *gin.Context) {
	userId, err := utils.ParseUserIDInToken(c)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
		return
	}

	data := api.UploadThesis{}
	if err := c.Bind(&data); err != nil {
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
		return
	}

	blob, err := c.FormFile("file")
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
		return
	}

	file := model.File{
		Name: blob.Filename,
		Blob: blob,
	}

	if err := file.GetPath(); err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionCreate, logger.ErrorBodyCreatePath, err)
		return
	}

	thesis, err := services.UploadThesis(userId, data.ThesisId, file)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionUpdate, logger.ErrorBodyUpdateThesis, err)
		return
	}

	if err := file.Save(c); err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionSave, logger.ErrorBodySaveThesis, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"thesisId": thesis.ID,
	})
}

func ThesisReviewRecord(c *gin.Context) {
	userId, err := utils.ParseUserIDInToken(c)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
		return
	}

	operationLog, err := services.QueryReviewRecord(userId)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": operationLog,
	})
}

func ThesisToBeReviewedList(c *gin.Context) {
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

	if user.Role != model.Admin {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionQuery, logger.ErrorBodyPermissions, fmt.Errorf("Administrator Permission Is Required"))
		return
	}

	result, err := services.QueryToBeReviewedThesisList()
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyToBeReviewedList, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func ThesisAllocation(c *gin.Context) {
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

	if user.Role != model.Admin {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionQuery, logger.ErrorBodyPermissions, fmt.Errorf(""))
		return
	}

	data := api.AllocationThesis{}
	if err := c.BindJSON(&data); err != nil {
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
		return
	}

	if err := services.UpdateAllocationThesis(data); err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionQuery, logger.ErrorBodyAssignmentThesis, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func ThesisRandomAllocation(c *gin.Context) {
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

	if user.Role != model.Admin {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionQuery, logger.ErrorBodyPermissions, fmt.Errorf(""))
		return
	}

	data := api.RandomAllocationThesis{}
	if err := c.BindJSON(&data); err != nil {
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
		return
	}

	services.UpdateThesisRandomAllocation(data)
}

func ThesisUnderReviewList(c *gin.Context) {
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
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionQuery, logger.ErrorBodyUnderReviewList, fmt.Errorf(""))
		return
	}

	result, err := services.QueryUnderReviewThesisList(user)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionQuery, logger.ErrorBodyUnderReviewList, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func ThesisDownload(c *gin.Context) {
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
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionQuery, logger.ErrorBodyUnderReviewList, fmt.Errorf(""))
		return
	}

	data := api.DownloadThesis{}
	if err := c.BindJSON(&data); err != nil {
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
		return
	}

	result, err := services.QueryThesisDownloadPath(data)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyThesisDownloadLink, err)
		return
	}

	c.File(result.FileURL)
}
