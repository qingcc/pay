package api

import (
	"encoding/json"
	common "github.com/qingcc/wechat/utils/ipayLinks/model"
	api_model "github.com/qingcc/wechat/utils/ipayLinks/model/api"
	tools2 "github.com/qingcc/wechat/utils/ipayLinks/tools"
	"io/ioutil"
	"net/http"
)

//同步取消
func CancelOrder(w http.ResponseWriter, r *http.Request) {
	rb, _ := ioutil.ReadAll(r.Body)
	req := new(api_model.CancelOrder)
	json.Unmarshal(rb, &req)
	req.Init()

	res := new(api_model.CancelResponse)
	s := common.DefaultNewSign(req)
	tools2.IpayLinksSendData(s, res, true)
	sign := &common.Sign{
		Sign: res.Sign,
		Data: res,
	}
	dealPay(sign, &res.CommonResponse, w)
}

//异步取消
func CancelOrderSync(w http.ResponseWriter, r *http.Request) {
	rb, _ := ioutil.ReadAll(r.Body)
	cancelReq := new(api_model.CancelResponse)
	json.Unmarshal(rb, &cancelReq)
	sign := &common.Sign{
		Sign: cancelReq.Sign,
		Data: cancelReq,
	}
	dealPay(sign, &cancelReq.CommonResponse, w)
}
