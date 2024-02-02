package common

// Reply Status and Message
const (
	ReplyStatusSuccess  = 0
	ReplyMessageSuccess = "成功"

	ReplyStatusBindRequestFailed  = 1
	ReplyMessageBindRequestFailed = "API参数错误"
	ReplyStatusMissingParam       = 2
	ReplyMessageMissingParam      = "缺失API参数"

	ReplyStatusCreateFailed  = 10
	ReplyMessageCreateFailed = "创建失败"

	ReplyStatusReadFailed  = 20
	ReplyMessageReadFailed = "读取数据失败"

	ReplyStatusUpdateFailed  = 30
	ReplyMessageUpdateFailed = "更新失败"

	ReplyStatusDeleteFailed  = 40
	ReplyMessageDeleteFailed = "删除失败"
)

// RecordStatus 0:初始状态，-1:删除
const (
	RecordStatusInit    = 0
	RecordStatusDeleted = -1
)
