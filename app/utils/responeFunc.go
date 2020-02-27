package utils

type ResponseInfo struct {
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Data    interface{} `json:"data"`
}

func Respone(errCode int, errMsg string, data interface{}) ResponseInfo {
	return ResponseInfo{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    data,
	}
}
