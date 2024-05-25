package api

type Login struct {
	Account  string `json:"account" binding:"required"`  // 账户
	Password string `json:"password" binding:"required"` // 密码
}

type Administrator struct {
	Name     string `json:"name" binding:"required"`
	College  string `json:"college" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Identity string `json:"identity" binding:"required"`
}

type Teacher struct {
	Name     string `json:"name" binding:"required"`
	College  string `json:"college" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Identity string `json:"identity" binding:"required"`
}

type Student struct {
	Name     string `json:"name" binding:"required"`
	College  string `json:"college" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Subject  string `json:"subject" binding:"required"`
	Class    string `json:"class" binding:"required"`
	Identity string `json:"identity" binding:"required"`
}

type TeacherList struct {
	Id      int    `json:"id" binding:"required"`      // 教师 ID
	Name    string `json:"name" binding:"required"`    // 姓名
	Phone   string `json:"phone" binding:"required"`   // 手机号
	Number  string `json:"number" binding:"required"`  // 工号
	College string `json:"college" binding:"required"` // 院系
}
