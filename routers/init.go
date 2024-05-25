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
			// user.POST("/editor", controller.UserEditor)
			// user.POST("/register", controller.UserRegister)
			user.POST("/logout", controller.UserLogout)
			user.POST("/teacherList", controller.UserTeacherList)
		}

		thesis := v1.Group("/thesis")
		{
			thesis.POST("/create", controller.ThesisCreate)
			thesis.POST("/upload", controller.ThesisUpload)
			thesis.POST("/reviewRecord", controller.ThesisReviewRecord)
			thesis.POST("/allotList", controller.ThesisToBeReviewedList)
			thesis.POST("/allocation", controller.ThesisAllocation)
			thesis.POST("/randomAllocation", controller.ThesisRandomAllocation)

			thesis.POST("/reviewList", controller.ThesisUnderReviewList)
			thesis.POST("/download", controller.ThesisDownload)
		}

		review := v1.Group("/review")
		{
			review.POST("upload")
		}
	}
}
