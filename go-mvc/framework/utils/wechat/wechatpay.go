package wechat

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"crypto/tls"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WechatPay struct {
	AppId         string
	MchId         string
	ApiKey        string
	ApiclientCert []byte
	ApiclientKey  []byte
}

func New(appId, mchId, apiKey string, apiclient_cert, apiclient_key []byte) (client *WechatPay) {
	client = &WechatPay{}
	client.AppId = appId
	client.MchId = mchId
	client.ApiKey = apiKey
	client.ApiclientCert = apiclient_cert
	client.ApiclientKey = apiclient_key
	return client
}

//wxpay计算签名
func GetSign(mReq map[string]interface{}, key string) (sign string) {

	sortedKeys := make([]string, 0)

	for k, _ := range mReq {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	var signStrings string

	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

//wxpay验证签名
func (this *WechatPay) VerifySign(needVerifyM map[string]interface{}, sign string) bool {
	delete(needVerifyM,"sign")
	signCalc := GetSign(needVerifyM, this.ApiKey)

	if sign == signCalc {
		return true
	}
	return false
}

func WithCertBytes(cert, key []byte) *http.Transport {
	tlsCert, err := tls.X509KeyPair(cert, key)

	if err != nil {
		return nil
	}
	conf := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	}
	trans := &http.Transport{
		TLSClientConfig: conf,
	}
	return trans
}

func randomNonceStr() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
