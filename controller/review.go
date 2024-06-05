package controller

import (
	"net/http"

	"github.com/ZEQUANR/zhulong/logger"
	"github.com/ZEQUANR/zhulong/model"
	"github.com/ZEQUANR/zhulong/model/api"
	"github.com/ZEQUANR/zhulong/services"
	"github.com/ZEQUANR/zhulong/utils"
	"github.com/gin-gonic/gin"
)

func ReviewUpload(c *gin.Context) {
	userId, err := utils.ParseUserIDInToken(c)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionParse, logger.ErrorBodyParseToken, err)
		return
	}

	data := api.UploadReview{}
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

	_, err = services.UploadReview(userId, data.ThesisId, file)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionUpdate, logger.ErrorBodyUpdateReview, err)
		return
	}

	if err := file.Save(c); err != nil {
		logger.CreateLog(c, logger.ErrorWhoServer, logger.ErrorActionSave, logger.ErrorBodySaveReview, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func ReviewSendBack(c *gin.Context) {
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

	data := api.SendBackThesis{}
	if err := c.Bind(&data); err != nil {
		logger.CreateLog(c, logger.ErrorWhoClient, logger.ErrorActionRead, logger.ErrorBodyParameters, err)
		return
	}

	_, err = services.RecordThesisSendBack(user, data)
	if err != nil {
		logger.CreateLog(c, logger.ErrorWhoDatabase, logger.ErrorActionQuery, logger.ErrorBodyReturnThesis, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
