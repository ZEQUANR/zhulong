package api

type CreateThesis struct {
	ChineseTitle  string `json:"chineseTitle" binding:"required"`  // 论文中文标题
	EnglishTitle  string `json:"englishTitle" binding:"required"`  // 论文英文标题
	Authors       string `json:"authors" binding:"required"`       // 论文作者
	Teachers      string `json:"teachers" binding:"required"`      // 论文指导教师
	FirstAdvance  string `json:"firstAdvance" binding:"required"`  // 第一创新点
	SecondAdvance string `json:"secondAdvance" binding:"required"` // 第二创新点
	ThirdAdvance  string `json:"thirdAdvance" binding:"required"`  // 第三创新点
	Drawback      string `json:"drawback" binding:"required"`      // 论文不足
}

type UploadThesis struct {
	ThesisId int `bson:"thesis_id" form:"thesis_id"` // 论文 ID
}
