package utils

type ResponseInfo struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessRespone(data interface{}) ResponseInfo {
	return ResponseInfo{
		Code: 1,
		Msg:  "success",
		Data: data,
	}
}

func FailRespone(code int64, msg string, data interface{}) ResponseInfo {
	return ResponseInfo{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
