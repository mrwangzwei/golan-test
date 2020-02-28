package utils

type ResponseInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Respone(code int, msg string, data interface{}) ResponseInfo {
	return ResponseInfo{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
