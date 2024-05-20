package logger

type ErrorWho string

const (
	ErrorWhoServer   ErrorWho = "Server"   // 服务端
	ErrorWhoClient   ErrorWho = "Client"   // 客户端
	ErrorWhoDatabase ErrorWho = "Database" // 数据库
)

type ErrorAction string

const (
	ErrorActionRead   ErrorAction = "Read"   // 读取
	ErrorActionQuery  ErrorAction = "Query"  // 查询
	ErrorActionUpdate ErrorAction = "Update" // 更新
	ErrorActionCreate ErrorAction = "Create" // 创建
	ErrorActionSave   ErrorAction = "Save"   // 保存
	ErrorActionParse  ErrorAction = "Parse"  // 解析
)

type ErrorBody string

const (
	ErrorBodyParameters         ErrorBody = "Invalid Request Parameters"                         // 无效的请求参数
	ErrorBodyRequestHeader      ErrorBody = "Invalid Request Header"                             // 无效的请求头
	ErrorBodyQueryingUser       ErrorBody = "Failed Querying User"                               // 查询用户失败
	ErrorBodyCreateToken        ErrorBody = "Failed To Create Token"                             // 创建 Token 失败
	ErrorBodyParseToken         ErrorBody = "Failed To Parse Token"                              // 解析 Token 失败
	ErrorBodyCreateThesis       ErrorBody = "Thesis Creation Failure"                            // 创建论文失败
	ErrorBodyCreatePath         ErrorBody = "Failed To Create Token"                             // 创建 Path 失败
	ErrorBodyUpdateThesis       ErrorBody = "Failed To Update The Thesis Information"            // 更新论文信息失败
	ErrorBodySaveThesis         ErrorBody = "Failed To Save The Thesis File"                     // 保存论文文件失败
	ErrorBodyPermissions        ErrorBody = "The Current User Has Insufficient Permissions"      // 当前用户权限不足
	ErrorBodyAssignmentThesis   ErrorBody = "Assignment Thesis Failure"                          // 分配论文失败
	ErrorBodyToBeReviewedList   ErrorBody = "Failed To Obtain The List Of Thesis To Be Reviewed" // 获取待评阅论文列表失败
	ErrorBodyUnderReviewList    ErrorBody = "Failed To Obtain A List Of Thesis Under Review"     // 获取评阅中论文列表失败
	ErrorBodyThesisDownloadLink ErrorBody = "Failed To Obtain The Thesis Download Link"          // 获取论文下载链接失败
)

// type ErrorCause struct {
// 	Content string
// 	Code    int
// }

// var (
// 	StatusBadRequest          = ErrorCause{Content: "Invalid request parameters", Code: http.StatusBadRequest}     // 400 无效的请求参数
// 	StatusUnauthorized        = ErrorCause{Content: "Authentication Failure", Code: http.StatusUnauthorized}       // 401 身份验证失败
// 	StatusInternalServerError = ErrorCause{Content: "Internal Server Error", Code: http.StatusInternalServerError} // 500 内部服务器错误
// )
