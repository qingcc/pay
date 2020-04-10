package api

import (
	"encoding/json"
	"github.com/qingcc/wechat/utils/ipayLinks/model"
	"github.com/qingcc/wechat/utils/ipayLinks/model/api"
	tools2 "github.com/qingcc/wechat/utils/ipayLinks/tools"
	"github.com/qingcc/yi/commobj"
	"io/ioutil"
	"net/http"
)

func PayAsyn()  {
	req := api_model.PayRequest{} //todo fill request param
	res := new(api_model.PayResponseAsyn)
	s := common.DefaultNewSign(req)
	tools2.IpayLinksSendData(s, res, true) //TODO 待完成 请求写入日志
	sign := &common.Sign{
		Sign: res.Sign,
		Data: res,
	}
	if sign.VerifyNotifySignData() {
		//verified todo something
	}else {
		//verify failed todo something
	}
}

//check response is error
func checkError(httpMessageResult commobj.HttpMessageResult, interfaceType string) (success bool) {
	//result.SplMessageLogs = append(result.SplMessageLogs, httpMessageResult)
	if httpMessageResult.Success {
	} else {
		return false
	}
	return true
}

//支付异步响应（对方使用post方法将数据异步推送过来，接收请求参数并响应
func PaySync(w http.ResponseWriter, r *http.Request)  {
	s, _ := ioutil.ReadAll(r.Body)
	payResSync := new(api_model.PayResponseSync)
	json.Unmarshal(s, &payResSync)
	sign := &common.Sign{
		Sign: payResSync.Sign,
		Data: payResSync,
	}
	if sign.VerifyNotifySignData() {
		//todo something 保存数据
	}else {
		//todo something
	}
	return
}
