package model

var OperationLogStatus = struct {
	UnderReview int // 待评阅
	OnReview    int // 评阅中
}{
	UnderReview: 1,
	OnReview:    2,
}

var OperationLogContext = struct {
	ThesisUpload   int // 论文上传
	ThesisReviewer int // 论文评阅
}{
	ThesisUpload:   1,
	ThesisReviewer: 2,
}
