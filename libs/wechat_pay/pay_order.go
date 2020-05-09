package wechat_pay

import (
	"bytes"
	"crypto/md5"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"self-test/app/utils"
	"strconv"
	"strings"
	"time"
)

const (
	payUri    = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	signType  = "MD5"
	tradeType = "JSAPI"
)

type JsApiPayConfig struct {
	Appid     string
	TimeStamp string
	NonceStr  string
	Package   string
	SignType  string
	PaySign   string
}

type unifiedRes struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Openid     string `xml:"openid"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
	ErrCodeDes string `xml:"err_code_des"`
}

type unifiedOrderParam struct {
	XMLName    xml.Name `xml:"xml"`
	Appid      string   `xml:"appid" json:"appid"`
	Body       string   `xml:"body" json:"body"`
	MchId      string   `xml:"mch_id" json:"mch_id"`
	NonceStr   string   `xml:"nonce_str" json:"nonce_str"`
	NotifyUrl  string   `xml:"notify_url" json:"notify_url"`
	Openid     string   `xml:"openid" json:"openid"`
	OutTradeNo string   `xml:"out_trade_no" json:"out_trade_no"`
	Sign       string   `xml:"sign" json:"sign"`
	SignType   string   `xml:"sign_type" json:"sign_type"`
	TotalFee   int64    `xml:"total_fee" json:"total_fee"`
	TradeType  string   `xml:"trade_type" json:"trade_type"`
}

func GetJsApiPayConfig(totalFree int64, openid, appid, mchId, mchKey, body, orderNo, notifyUrl string) (res JsApiPayConfig, err error) {
	nonceStr := utils.RandString(10)
	op := unifiedOrderParam{
		Openid:     openid,
		Appid:      appid,
		MchId:      mchId,
		NonceStr:   nonceStr,
		Sign:       "",
		SignType:   signType,
		Body:       body,
		OutTradeNo: orderNo,
		TotalFee:   totalFree,
		NotifyUrl:  notifyUrl,
		TradeType:  tradeType,
	}
	op.Sign, err = sign(&op, mchKey)
	if err != nil {
		return
	}
	resp, err := unifiedOrder(&op)
	if err != nil {
		return
	}
	if resp.ReturnCode != "SUCCESS" || resp.ResultCode != "SUCCESS" {
		return res, errors.New(resp.ReturnMsg + resp.ErrCodeDes)
	}
	res.Appid = resp.Appid
	res.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
	res.NonceStr = nonceStr
	res.Package = resp.PrepayId
	res.SignType = "MD5"
	res.PaySign = jsapiSign(&res, mchKey)
	return res, nil
}

func jsapiSign(p *JsApiPayConfig, mchKey string) string {
	v := url.Values{}
	v.Add("appId", p.Appid)
	v.Add("nonceStr", p.NonceStr)
	v.Add("package", "prepay_id="+p.Package)
	v.Add("signType", p.SignType)
	v.Add("timeStamp", p.TimeStamp)
	r, _ := url.QueryUnescape(v.Encode())
	r += "&key=" + mchKey
	has := md5.Sum([]byte(r))
	sign := fmt.Sprintf("%x", has)
	sign = strings.ToUpper(sign)
	return sign
}

func sign(op *unifiedOrderParam, mchKey string) (sign string, err error) {
	v := url.Values{}
	v.Add("appid", op.Appid)
	v.Add("body", op.Body)
	v.Add("mch_id", op.MchId)
	v.Add("nonce_str", op.NonceStr)
	v.Add("notify_url", op.NotifyUrl)
	v.Add("openid", op.Openid)
	v.Add("out_trade_no", op.OutTradeNo)
	v.Add("sign_type", op.SignType)
	v.Add("total_fee", strconv.FormatInt(op.TotalFee, 10))
	v.Add("trade_type", op.TradeType)
	r, _ := url.QueryUnescape(v.Encode())
	r += "&key=" + mchKey
	has := md5.Sum([]byte(r))
	sign = fmt.Sprintf("%x", has)
	sign = strings.ToUpper(sign)
	return
}

func unifiedOrder(op *unifiedOrderParam) (res unifiedRes, err error) {
	var xmlStr []byte
	if xmlStr, err = xml.Marshal(op); err != nil {
		return
	}
	resp, err := http.Post(payUri,
		"application/x-www-form-urlencoded",
		bytes.NewBuffer(xmlStr))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err = xml.Unmarshal(body, &res); err != nil {
		return
	}
	return
}
