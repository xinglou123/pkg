package utils

const (
	MSG_OK       = 0  // 成功
	MSG_ERR      = -1 // 错误
	MSG_REDIRECT = -2 // 重定向
	MSG_AUTH     = -3 // 登录认证
)

// Response represents HTTP response body.
type Response struct {
	Code int         `json:"code"` // return code, 0 for succ
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}

// NewResponse creates a result with Code=0, Msg="", Data=nil.
func NewResponse() *Response {
	return &Response{
		Code: MSG_OK,
		Msg:  "",
		Data: nil,
	}
}
