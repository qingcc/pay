package api

import (
	"encoding/json"
	"github.com/qingcc/wechat/http/model"
	"github.com/qingcc/wechat/utils"
	common "github.com/qingcc/wechat/utils/ipayLinks/model"
	api_model "github.com/qingcc/wechat/utils/ipayLinks/model/api"
	"github.com/qingcc/wechat/utils/ipayLinks/tools"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
* 使用token消费
 */

func TokenConsume(w http.ResponseWriter, r *http.Request) {
	req := new(api_model.TokenConsumeRequest)
	by, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(by, req)
	req.CarrierId = common.CarrierId
	req.Init()
	s := common.DefaultNewSign(req)

	res := new(api_model.TokenConsumeResponse)
	tools.IpayLinksSendData(s, res, true)
	sign := &common.Sign{
		Sign: res.Sign,
		Data: res,
	}
	dealTokenConsume(sign, &res.CommonResponse, w)
}

func TokenConsumeSync(w http.ResponseWriter, r *http.Request) {
	res := new(api_model.TokenConsumeResponse)
	by, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(by, res)

	sign := &common.Sign{
		Sign: res.Sign,
		Data: res,
	}
	dealTokenConsume(sign, &res.CommonResponse, w)
}

//处理下单响应（同步或异步） 同步，响应给调用支付接口的上游服务； 异步，响应给ipayLinks
func dealTokenConsume(sign *common.Sign, res *api_model.CommonResponse, w http.ResponseWriter) {
	if sign.VerifyNotifySignData() {
		//verified
		orderId, _ := strconv.Atoi(res.TransId)
		status := api_model.OrderStatus(res.Status)
		o := common.Order{OrderId: orderId, Status: status}
		go o.Update("status", "order_id", o.OrderId)

		w.Write(utils.ApiRetrun(http.StatusOK, "", model.ApiOrderResponse{OrderStatus: status})) //
	} else {
		//verify failed
		w.Write(utils.ApiRetrun(http.StatusInternalServerError, "服务器内部错误", nil))
	}
}
