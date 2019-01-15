package utils

// Response represents HTTP response body.
type Response struct {
	Code int         `json:"code"` // return code, 0 for succ
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}

// NewResponse creates a result with Code=0, Msg="", Data=nil.
func NewResponse() *Response {
	return &Response{
		Code: 200,
		Msg:  "",
		Data: nil,
	}
}
