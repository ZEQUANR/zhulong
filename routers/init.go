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
			user.POST("/register", controller.UserRegister)
			user.POST("/info", controller.UserInfo)
			user.POST("/logout", controller.UserLogin)
		}
	}
}
