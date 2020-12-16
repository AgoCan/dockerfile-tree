package response

// 错误码
const (
	ErrCodeSuccess   = 0
	ErrCodeParameter = 1001
	ErrSQLList       = 2001
)

func getMessage(code int) (message string) {
	var codeMsgMap map[int]string
	codeMsgMap = make(map[int]string, 1024)

	codeMsgMap[ErrCodeSuccess] = "success"
	codeMsgMap[ErrCodeParameter] = "参数错误"
	codeMsgMap[ErrSQLList] = "查询错误"

	message, ok := codeMsgMap[code]
	if !ok {
		message = "未知错误"
	}
	return message
}