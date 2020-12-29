package serializer

var codeMsgMap map[int]string

// 错误码
const (
	CodeSuccess = 0
	// 4 开头的是前端操作问题
	ErrCodeParameter = 41001
	// 5 开头是后端问题
	ErrSQL         = 52001
	ErrSQLExist    = 52002
	ErrSQLNotExist = 52003
)

func init() {
	codeMsgMap = make(map[int]string, 1024)
	codeMsgMap[CodeSuccess] = "success"
	codeMsgMap[ErrCodeParameter] = "参数错误"
	codeMsgMap[ErrSQL] = "sql错误"
	codeMsgMap[ErrSQLExist] = ": 已存在"
	codeMsgMap[ErrSQLNotExist] = "数据不存在"
}

// GetMessage 获取错误信息
func GetMessage(code int) (message string) {
	message, ok := codeMsgMap[code]
	if !ok {
		message = "未知错误"
	}
	return message
}
