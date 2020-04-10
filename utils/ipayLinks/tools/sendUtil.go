package tools

import (
	"encoding/json"
	model2 "github.com/qingcc/wechat/utils/ipayLinks/model"
	"github.com/qingcc/yi/commobj"
	"github.com/qingcc/yi/utils/httputils"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	httpService = httputils.ServiceHttpClient{}
)

func IpayLinksSendData(s *model2.Sign, res interface{}, storage bool) (httpMessageResult commobj.HttpMessageResult) {
	httpMessageResult.Start = time.Now()
	resultMessageArr := make([]string, 0)

	s.Signature()
	url := s.Domain
	var(
		code int
		message string
		data []byte
	)
	if s.Method == http.MethodPost {
		jsonReq, err := json.Marshal(s.Data)
		if err != nil {
			resultMessageArr = append(resultMessageArr, "[Marshal Request]fail"+err.Error())
			return
		}
		resultMessageArr = append(resultMessageArr, "[Marshal Request]Success")
		httpMessageResult.Req = string(jsonReq)
		code, message, data = httpService.SendJson(url, jsonReq, s.Header, 1, res)
	}else if s.Method == http.MethodGet {
		url += "?" + s.Param
		code, message, data = httpService.GetJson(url, s.Header, 1, res)
	}

	httpMessageResult.ResultCode = code
	httpMessageResult.SplResultCode = code
	httpMessageResult.Url = url

	httpMessageResult.Res = string(data)
	if storage {
		log.Printf("请求对象==>%v\n", httpMessageResult.Req)
		log.Printf("返回对象==>%v\n", httpMessageResult.Res)
	}
	if code == commobj.SUCCESS {
		httpMessageResult.Success = true
		resultMessageArr = append(resultMessageArr, "[SendJson Message]Success")
	} else {
		resultMessageArr = append(resultMessageArr, "[SendJson Message]fail"+message)
	}
	httpMessageResult.ResultMessage = strings.Join(resultMessageArr, ",")
	httpMessageResult.SplResultMessage = strings.Join(resultMessageArr, ",")
	httpMessageResult.End = time.Now()
	return
}

