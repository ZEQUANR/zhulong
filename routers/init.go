package routers

import (
	"github.com/ZEQUANR/zhulong/controller"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/login", controller.UserLogin)
			user.POST("/info", controller.UserInfo)
			user.POST("/editor", controller.UserEditor)
			user.POST("/register", controller.UserRegister)
			user.POST("/logout", controller.UserLogout)
			user.POST("/redact")
		}

		thesis := v1.Group("/thesis")
		{
			thesis.POST("/Uploading")
			thesis.POST("/download")
		}
	}
}
