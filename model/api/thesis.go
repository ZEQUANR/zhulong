package api

type CreateThesis struct {
	ThesisTitle string `json:"thesis_title" binding:"required"`
}

type UploadThesis struct {
	ThesisId int `bson:"thesis_id" form:"thesis_id"`
}
