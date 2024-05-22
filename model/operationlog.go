package model

var OperationLogStatus = struct {
	UnderReview   int // 审核中
	AlreadyPassed int // 已通过
}{
	UnderReview:   1,
	AlreadyPassed: 2,
}

var OperationLogContext = struct {
	ThesisUpload   int // 论文上传
	ThesisReviewer int // 论文评阅
}{
	ThesisUpload:   1,
	ThesisReviewer: 2,
}
