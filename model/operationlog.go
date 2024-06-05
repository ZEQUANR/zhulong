package model

var OperationLogStatus = struct {
	SendBack        int // 论文退回
	UnderReview     int // 论文提交
	OnReview        int // 评阅中
	ReviewCompleted int // 评阅完成
}{
	SendBack:        0,
	UnderReview:     1,
	OnReview:        2,
	ReviewCompleted: 3,
}

var OperationLogContext = struct {
	ThesisUpload    string // 上传论文
	ThesisReviewer  string // 进行评阅
	UploadTheReview string // 上传评阅书
}{
	ThesisUpload:    "上传论文",
	ThesisReviewer:  "进行评阅",
	UploadTheReview: "上传评阅书",
}
