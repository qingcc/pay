package common

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/qingcc/yi/utils"
	"io"
	"log"
	"net/url"
	"sort"
	"strings"
)

var(
	test_url = "https://mapi1.uat.ipaylinks.com/mapi/OpenAPI.do" //联调
	inner_url = "https://mapi-aw.ipaylinks.com/mapi/OpenAPI.do"//product 服务器在国内
	fore_url = "https://mapi.ipaylinks.com/mapi/OpenAPI.do" //product 服务器在国外
	signKey = "md5SignKey"
	privateKeyFile = "static/keys/private.pem"
	publicKeyFile = "static/keys/public.pem"
	header = map[string]string{}
	selectedSignType = signTypeRsa
)

var (
	signTypeMd5 = "MD5"
	signTypeRsa = "rsa"
)
type Sign struct {
	Sign string `json:"sign"` //存储签名之后生成的签名字符串
	SignType string `json:"signType"`//签名类型（eg. MD5 or RSA)
	SignData []byte `json:"-"` //需要使用私钥签名的数据
	DecryptData []byte `json:"-"`//需要使用私钥解密的数据
	Method string `json:"-"` //请求方式（eg. http.MethodPost or http.MethodGet)
	Domain string    `json:"-"` //请求域名
	Header map[string]string `json:"-"` //请求header
	Data interface{} //其他的请求参数
	Param string     `json:"-"`//为GET请求时，生成的已经签名带参数的uri
}

func DefaultNewSign(data interface{}) *Sign {
	return NewSign(selectedSignType, "POST", 2, header, data)
}

func NewSign(signType, method string, domainType int, h map[string]string, data interface{}) *Sign {
	domain := fore_url
	switch domainType {
	case 0:
		domain = test_url
	case 1:
		domain = inner_url
	}
	return &Sign{
		SignType:    signType,
		Method:      method,
		Domain:      domain,
		Header:      h,
		Data:		 data,
	}
}


// 签名
func (s *Sign)Signature() (err error) {
	dataMap, signSlice := make(map[string]string), make([]string, 0)
	jsonData, _ := json.Marshal(s.Data)
	json.Unmarshal(jsonData, &dataMap)
	for k, v := range dataMap {
		if strings.TrimSpace(v) != "" {
			signSlice = append(signSlice, k + "=" + v)
		}
	}
	sort.Strings(signSlice)

	s.SignData = []byte(strings.Join(signSlice, "&"))
	err = s.signContent()

	log.Println("sign before:", strings.Join(signSlice, "&"))
	signSlice = append(signSlice, "sign="+ s.Sign)
	s.Param = url.QueryEscape(strings.Join(signSlice, "&"))
	log.Println("sign after:", s.Param)
	return
}

func (s *Sign)signContent() (err error) {
	switch s.SignType {
	case signTypeMd5:
		err = s.signMd5()
	case signTypeRsa:
		utils.LoadKeyFile(privateKeyFile, publicKeyFile)
		err = s.signRsa()
	}
	return
}

func (s *Sign)signMd5() (err error) {
	signMd5 := string(s.SignData) + "signKey" + signKey
	log.Printf("待签名数据： %v", signMd5)
	newMd5 := md5.New()
	_, err = io.WriteString(newMd5, signMd5)
	s.Sign = fmt.Sprintf("%x", newMd5.Sum(nil))
	log.Printf("生成签名: %v", s.Sign)
	return
}

//私钥签名（公钥验证）
func (s *Sign)signRsa() error {
	signedData, err := utils.RsaSign(s.SignData)
	s.Sign = string(signedData)
	return err
}

//私钥解密（公钥加密）
func (s *Sign)decryptRsa() error {
	decryptedData, err := utils.RsaDecrypt(s.DecryptData)
	s.Sign = string(decryptedData)
	return err
}

func struct2Slice(data interface{}) []string {
	dataMap, dataSlice := make(map[string]interface{}), make([]string, 0)
	dataJson, _ := json.Marshal(&data)
	json.Unmarshal(dataJson, &dataMap)

	for key, value := range dataMap {
		if key == "sign" {
			continue
		}
		dataSlice = append(dataSlice, fmt.Sprintf("%s=%v", key, value))
	}
	sort.Strings(dataSlice)
	return dataSlice
}

//同步响应/异步通知 签名验证
func (s *Sign)VerifyNotifySignData() (verified bool) {
	sign := s.Sign
	dataSlice := struct2Slice(s.Data)
	switch s.SignType {
	case signTypeMd5:
		s.signMd5()
	case signTypeRsa:
		for _, v := range dataSlice {
			sign += v
		}
		s.Sign += utils.AesDecrypter(s.Sign)
	}
	if sign == s.Sign { //检验成功
		verified = true
	}
	return
}