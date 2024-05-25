package api

import "time"

type CreateThesis struct {
	ChineseTitle  string   `json:"chineseTitle" binding:"required"`  // 论文中文标题
	EnglishTitle  string   `json:"englishTitle" binding:"required"`  // 论文英文标题
	Authors       []string `json:"authors" binding:"required"`       // 论文作者
	Teachers      []string `json:"teachers" binding:"required"`      // 论文指导教师
	FirstAdvance  string   `json:"firstAdvance" binding:"required"`  // 第一创新点
	SecondAdvance string   `json:"secondAdvance" binding:"required"` // 第二创新点
	ThirdAdvance  string   `json:"thirdAdvance" binding:"required"`  // 第三创新点
	Drawback      string   `json:"drawback" binding:"required"`      // 论文不足
}

type UploadThesis struct {
	ThesisId int `form:"thesisId" binding:"required"` // 论文 ID
}

type DownloadThesis struct {
	ThesisId int `json:"thesisId" binding:"required"` // 论文 ID
}

type ToBeReviewedThesisList struct {
	Id           int       `json:"id" binding:"required"`           // 论文 ID
	Role         int       `json:"role" binding:"required"`         // 权限
	Name         string    `json:"name" binding:"required"`         // 姓名
	Number       string    `json:"number" binding:"required"`       // 学号
	College      string    `json:"college" binding:"required"`      // 院系
	FileState    int       `json:"fileState" binding:"required"`    // 论文状态
	ChineseTitle string    `json:"chineseTitle" binding:"required"` // 论文中文标题
	EnglishTitle string    `json:"englishTitle" binding:"required"` // 论文英文标题
	UploadTime   time.Time `json:"uploadTime" binding:"required"`   // 论文上传时间
}

type AllocationThesis struct {
	ThesisID  []int `json:"thesisIds" binding:"required"` // 论文 ID
	TeacherID int   `json:"teacherId" binding:"required"` // 老师 ID
}

type RandomAllocationThesis struct {
	ThesisID []int `json:"thesisIds" binding:"required"` // 论文 ID
}
