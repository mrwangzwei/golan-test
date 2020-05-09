package wechat_pay

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"self-test/app/utils"
	"self-test/config"
	"strconv"
	"strings"
)

const (
	transfersUri = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
	checkName    = "NO_CHECK"
)

//付款订单(设备号和ip可以不用)
type WithdrawOrder struct {
	XMLName        xml.Name `xml:"xml"`
	MchAppid       string   `xml:"mch_appid"`
	Mchid          string   `xml:"mchid"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	Openid         string   `xml:"openid"`
	CheckName      string   `xml:"check_name"`
	Amount         int      `xml:"amount"`
	Desc           string   `xml:"desc"`
}

//付款订单结果
type WithdrawResult struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	ResultCode     string `xml:"result_code"`
	MchAppid       string `xml:"mch_appid"`
	Mchid          string `xml:"mchid"`
	NonceStr       string `xml:"nonce_str"`
	ErrCode        string `xml:"err_code"`
	ErrCodeDes     string `xml:"err_code_des"`
	PaymentNo      string `xml:"payment_no"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	PaymentTime    string `xml:"payment_time"`
}

func WithdrawOrderTo(appid, openid, orderNo, desc string, amount int) (res WithdrawResult, err error) {
	order := WithdrawOrder{
		MchAppid:       appid,
		Mchid:          config.Conf.WechatPayConfig[appid].AppMchId,
		NonceStr:       utils.RandString(10),
		Sign:           "",
		PartnerTradeNo: orderNo,
		Openid:         openid,
		CheckName:      checkName,
		Amount:         amount,
		Desc:           desc,
	}
	signStr := md5WithdrawOrder(order, config.Conf.WechatPayConfig[appid].AppMchKey)
	order.Sign = signStr
	var xmlBody []byte
	xmlBody, err = xml.Marshal(order)
	if err != nil {
		return
	}
	var bodyByte []byte
	bodyByte, err = securePost(appid, transfersUri, xmlBody)
	if err != nil {
		return
	}
	if err = xml.Unmarshal(bodyByte, &res); err != nil {
		return
	}
	if res.ReturnCode != "SUCCESS" || res.ResultCode != "SUCCESS" {
		err = errors.New(res.ReturnMsg + res.ErrCodeDes)
		return
	}
	return
}

//md5签名
func md5WithdrawOrder(order WithdrawOrder, mchKey string) string {
	o := url.Values{}
	o.Add("mch_appid", order.MchAppid)
	o.Add("mchid", order.Mchid)
	o.Add("partner_trade_no", order.PartnerTradeNo)
	o.Add("check_name", order.CheckName)
	o.Add("amount", strconv.Itoa(order.Amount))
	o.Add("desc", order.Desc)
	o.Add("nonce_str", order.NonceStr)
	o.Add("openid", order.Openid)
	r, _ := url.QueryUnescape(o.Encode())
	has := md5.Sum([]byte(r + "&key=" + mchKey))
	sign := fmt.Sprintf("%x", has)
	sign = strings.ToUpper(sign)
	return sign
}

func getTLSConfig(appid string) (*tls.Config, error) {
	var _tlsConfig *tls.Config
	// load cert
	cert, err := tls.LoadX509KeyPair(config.Conf.WechatPayConfig[appid].AppMchCertPath, config.Conf.WechatPayConfig[appid].AppMchCertKeyPath)
	if err != nil {
		return nil, err
	}
	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	return _tlsConfig, nil
}

//携带ca证书的安全请求
func securePost(appid, url string, xmlContent []byte) ([]byte, error) {
	tlsConfig, err := getTLSConfig(appid)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(
		url,
		"application/xml",
		bytes.NewBuffer(xmlContent))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyByte, nil
}
