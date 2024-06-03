package api

type UploadReview struct {
	ThesisId int `form:"thesisId" binding:"required"` // 论文 ID
}
