package wechat_pay

import (
	"crypto/md5"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
)

func VerifyData(param []byte, mchKey string) (map[string]string, error) {
	paramData := make(map[string]string)
	err := xml.Unmarshal(param, (*ToMap)(&paramData))
	if err != nil {
		return nil, err
	}
	sign, err := backSign(paramData, mchKey)
	if paramData["sign"] != sign {
		return nil, errors.New("sign is error")
	}
	return paramData, err
}

func backSign(param map[string]string, mchKey string) (sign string, err error) {
	var mapkeys []string
	for key, _ := range param {
		if key != "sign" {
			mapkeys = append(mapkeys, key)
		}
	}
	sort.Strings(mapkeys)
	var signStr string
	for index, item := range mapkeys {
		if item != "sign" {
			if index != 0 {
				signStr += "&"
			}
			signStr += item + "=" + param[item]
		}
	}
	signStr += "&key=" + mchKey
	has := md5.Sum([]byte(signStr))
	sign = fmt.Sprintf("%x", has)
	sign = strings.ToUpper(sign)
	return
}

func ResponSuccessData(c *gin.Context) {
	str := "<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>"
	_, _ = c.Writer.WriteString(str)
}
