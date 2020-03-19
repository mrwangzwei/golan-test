package common

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadAndValidate(ioRead io.ReadCloser, requestParams RequestParams) (err error) {
	bytes, err := ioutil.ReadAll(ioRead)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, requestParams); err != nil{
		return err
	}
	// 校验参数
	return requestParams.Validator()
}

type RequestParams interface {
	Validator() error
}
