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
		"thesis_id": thesis.ID,
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
		"thesis_id": thesis.ID,
	})
}
