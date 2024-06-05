package api

type UploadReview struct {
	ThesisId int `form:"thesisId" binding:"required"` // 论文 ID
}

type SendBackThesis struct {
	ThesisId int    `form:"thesisId" binding:"required"` // 论文 ID
	Reason   string `form:"reason" binding:"required"`   // 原因
}
