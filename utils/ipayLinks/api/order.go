package api

import (
	"encoding/json"
	"github.com/qingcc/wechat/http/model"
	"github.com/qingcc/wechat/utils"
	"github.com/qingcc/wechat/utils/ipayLinks/model"
	"github.com/qingcc/wechat/utils/ipayLinks/model/api"
	tools2 "github.com/qingcc/wechat/utils/ipayLinks/tools"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
* 下单消费
 */

func PayAsyn(w http.ResponseWriter, r *http.Request) {
	rb, _ := ioutil.ReadAll(r.Body)
	payr := new(api_model.PayCreateRequest)
	json.Unmarshal(rb, &payr)
	payr.CarrierId = common.CarrierId
	payr.Init()
	res := new(api_model.PayResponse)
	s := common.DefaultNewSign(payr)
	tools2.IpayLinksSendData(s, res, true)
	sign := &common.Sign{
		Sign: res.Sign,
		Data: res,
	}
	dealPay(sign, &res.CommonResponse, w)
}

//支付异步响应（对方使用post方法将数据异步推送过来，接收请求参数并响应
func PaySync(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body)
	payRes := new(api_model.PayResponse)
	json.Unmarshal(s, &payRes)
	sign := &common.Sign{
		Sign: payRes.Sign,
		Data: payRes,
	}
	dealPay(sign, &payRes.CommonResponse, w)
}

//处理下单响应（同步或异步）
func dealPay(sign *common.Sign, res *api_model.CommonResponse, w http.ResponseWriter) {
	if sign.VerifyNotifySignData() {
		//verified
		orderId, _ := strconv.Atoi(res.TransId)
		status := api_model.OrderStatus(res.Status)
		o := common.Order{OrderId: orderId, Status: status}
		go o.Update("status", "order_id", o.OrderId)

		w.Write(utils.ApiRetrun(http.StatusOK, "", model.ApiOrderResponse{OrderStatus: status}))
	} else {
		//verify failed
		w.Write(utils.ApiRetrun(http.StatusInternalServerError, "服务器内部错误", nil))
	}
}
