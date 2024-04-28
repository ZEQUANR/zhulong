package logger

import (
	"fmt"
	"net/http"
)

type ErrorGroup string
type ErrorWho string

const (
	ErrorGroupWarning ErrorGroup = "warning" // 警告
	ErrorGroupError   ErrorGroup = "error"   // 错误

	ErrorWhoServer   ErrorWho = "server"   // 服务端
	ErrorWhoClient   ErrorWho = "client"   // 客户端
	ErrorWhoDatabase ErrorWho = "database" // 数据库
)

type ErrorType struct {
	Content string
	Code    int
}

var (
	ErrorTypeParameter     = ErrorType{Content: "parameter", Code: http.StatusBadRequest}                   // 参数
	ErrorTypePermissions   = ErrorType{Content: "permissions", Code: http.StatusUnauthorized}               // 权限
	ErrorTypeServer        = ErrorType{Content: "server", Code: http.StatusInternalServerError}             // 服务
	ErrorTypeCreate        = ErrorType{Content: "create", Code: http.StatusInternalServerError}             // 创建
	ErrorTypeDelete        = ErrorType{Content: "delete", Code: http.StatusInternalServerError}             // 删除
	ErrorTypeUpdate        = ErrorType{Content: "update", Code: http.StatusInternalServerError}             // 更新
	ErrorTypeQuery         = ErrorType{Content: "query", Code: http.StatusInternalServerError}              // 查询
	ErrorTypeNotFound      = ErrorType{Content: "not found", Code: http.StatusRequestedRangeNotSatisfiable} // 未找到
	ErrorTypeAlreadyExists = ErrorType{Content: "already exists", Code: http.StatusForbidden}               // 已存在
	ErrorTypeInsufficient  = ErrorType{Content: "insufficient", Code: http.StatusNotAcceptable}             // 不足
	ErrorTypeStatus        = ErrorType{Content: "status", Code: http.StatusPreconditionFailed}              // 状态
)

type ErrorBody string

const (
	ErrorBodyDataFormat      ErrorBody = "data format"       // 数据格式
	ErrorBodyDatabase        ErrorBody = "database"          // 数据库
	ErrorBodyToken           ErrorBody = "token"             // token
	ErrorBodyMobile          ErrorBody = "mobile"            // 电话
	ErrorBodyRole            ErrorBody = "role"              // 角色
	ErrorBodyDepartment      ErrorBody = "department"        // 部门
	ErrorBodyClinic          ErrorBody = "clinic"            // 诊室
	ErrorBodyOperatorID      ErrorBody = "operator id"       // 操作者 id
	ErrorBodyDoctorID        ErrorBody = "doctor id"         // 医生 id
	ErrorBodyDoctor          ErrorBody = "doctor"            // 医生
	ErrorIDCode              ErrorBody = "id code"           //身份证号
	ErrorBodyPatientID       ErrorBody = "patient id"        // 患者 id
	ErrorBodyPatientExtra    ErrorBody = "patient extra"     // 患者拓展信息
	ErrorBodyPatient         ErrorBody = "patient"           // 患者
	ErrorPrompterInfo        ErrorBody = "prompter info"     // 推荐人信息
	ErrorBodyMember          ErrorBody = "member"            // 会员
	ErrorBodyPassword        ErrorBody = "password"          // 密码
	ErrorBodyPayChannel      ErrorBody = "pay channel"       // 充值渠道
	ErrorBodyMail            ErrorBody = "mail"              // 邮寄信息
	ErrorBodyScheduler       ErrorBody = "scheduler"         // 排班
	ErrorBodySchedulerPeriod ErrorBody = "scheduler period"  // 排班时段
	ErrorBodyMedicineID      ErrorBody = "medicine id"       // 药品 id
	ErrorBodyMedicine        ErrorBody = "medicine"          // 药品
	ErrorBodyCommodity       ErrorBody = "commodity"         // 商品
	ErrorBodyCommodityID     ErrorBody = "commodity id"      // 商品 id
	ErrorBodyCommodityType   ErrorBody = "commodity type"    // 商品类型
	ErrorBodyConflictID      ErrorBody = "conflict id"       // 相冲药物id
	ErrorBodyImageID         ErrorBody = "image id"          // 图片 id
	ErrorBodyImage           ErrorBody = "image"             // 图片
	ErrorBodyDisease         ErrorBody = "disease"           // 疾病
	ErrorBodySymptom         ErrorBody = "symptom"           // 症状
	ErrorBodyRegistered      ErrorBody = "registered"        // 挂号
	ErrorBodyRegisteredFee   ErrorBody = "registered fee"    // 挂号费
	ErrorBodyRegisteredInfo  ErrorBody = "registered info"   // 挂号信息
	ErrorBodyRegisteredType  ErrorBody = "registered type"   // 挂号类型
	ErrorBodyCallNumber      ErrorBody = "call number"       // 叫号
	ErrorBodyCase            ErrorBody = "case"              // 病例
	ErrorBodyPrescription    ErrorBody = "prescription"      // 处方
	ErrorBodySample          ErrorBody = "sample"            // 留样
	ErrorBodyTemplate        ErrorBody = "template"          // 处方模板
	ErrorBodyRecharge        ErrorBody = "recharge"          // 充值
	ErrorBodyRefund          ErrorBody = "refund"            // 退费
	ErrorBodyMoney           ErrorBody = "money"             // 金额
	ErrorBodyGiftMoney       ErrorBody = "gift money"        // 赠送金额
	ErrorBodyBalance         ErrorBody = "balance"           // 余额
	ErrorBodyOrderID         ErrorBody = "order id"          // 订单 id
	ErrorBodyOrderStatus     ErrorBody = "order status"      // 订单状态
	ErrorBodyOrder           ErrorBody = "order"             // 订单
	ErrorBodyPaymentMethod   ErrorBody = "payment method"    // 支付方式
	ErrorBodyFileID          ErrorBody = "file id"           // 文件 id
	ErrorBodyFile            ErrorBody = "file"              // 文件
	ErrorBodyLexiconWords    ErrorBody = "lexicon words"     // 词库
	ErrorBodyInventory       ErrorBody = "inventory"         // 库存
	ErrorBodyBatch           ErrorBody = "batch"             // 批次
	ErrorBodyBatchInfo       ErrorBody = "update batch info" // 批量更新信息
	ErrorBodyTag             ErrorBody = "tag"               // 标签
	ErrorBodyBlackList       ErrorBody = "black list"        // 黑名单
	ErrorBodySetting         ErrorBody = "setting info"      // 设置信息
)

func newError(group ErrorGroup, who ErrorWho, errorType ErrorType, body ErrorBody, detail string) error {
	return fmt.Errorf(fmt.Sprintf("%s|%s|%s|%s|[%s]", group, who, errorType.Content, body, detail))
}
