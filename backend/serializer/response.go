package serializer

// Response 回调时的固定内容
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error 回调错误信息
func Error(code int) Response {
	return Response{
		Code:    code,
		Message: GetMessage(code),
	}

}

// ErrorStr 回调错误信息,自带错误信息
func ErrorStr(code int, str string) Response {
	return Response{
		Code:    code,
		Message: str,
	}

}

// Success 回调正确信息
func Success(data interface{}) Response {
	return Response{
		Code:    CodeSuccess,
		Message: GetMessage(CodeSuccess),
		Data:    data,
	}
}
