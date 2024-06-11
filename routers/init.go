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
			user.POST("/logout", controller.UserLogout)
			user.POST("/teacherList", controller.UserTeacherList)

			register := user.Group("/register")
			{
				register.POST("admin", controller.UserRegisterAdmin)
				register.POST("teacher", controller.UserRegisterTeacher)
				register.POST("student", controller.UserRegisterStudent)
			}

			editor := user.Group("editor")
			{
				editor.POST("admin", controller.UserEditorAdmin)
				editor.POST("teacher", controller.UserEditorTeacher)
				editor.POST("student", controller.UserEditorStudent)
				editor.POST("password", controller.UserEditorPassword)
			}
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
			review.POST("upload", controller.ReviewUpload)
			review.POST("sendBack", controller.ReviewSendBack)
		}
	}
}
