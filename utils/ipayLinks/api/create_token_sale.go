package api

import (
	"encoding/json"
	"github.com/qingcc/wechat/utils"
	common "github.com/qingcc/wechat/utils/ipayLinks/model"
	api_model "github.com/qingcc/wechat/utils/ipayLinks/model/api"
	"github.com/qingcc/wechat/utils/ipayLinks/service"
	"github.com/qingcc/wechat/utils/ipayLinks/tools"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
* 消费者进行首次支付时，商户可通过<create_token_sale 接口>发起一笔消费交易。消费
* 交易成功后，iPayLinks 会为该消费者该卡片创建 token，并将 token 返回给商户
 */

func CreateTokenSale(w http.ResponseWriter, r *http.Request) {
	rb, _ := ioutil.ReadAll(r.Body)
	req := new(api_model.CreateTokenSaleRequest)
	json.Unmarshal(rb, &req)
	req.Init()

	res := new(api_model.CreateTokenSaleResponse)
	s := common.DefaultNewSign(req)
	tools.IpayLinksSendData(s, res, true)
	sign := &common.Sign{
		Sign: res.Sign,
		Data: res,
	}
	dealCreateTokenSaleResponse(sign, res, w)
}

//异步响应
func CreateTokenSaleSync(w http.ResponseWriter, r *http.Request) {
	rb, _ := ioutil.ReadAll(r.Body)
	cancelReq := new(api_model.CreateTokenSaleResponse)
	json.Unmarshal(rb, &cancelReq)
	sign := &common.Sign{
		Sign: cancelReq.Sign,
		Data: cancelReq,
	}
	dealCreateTokenSaleResponse(sign, cancelReq, w)
}

func dealCreateTokenSaleResponse(sign *common.Sign, res *api_model.CreateTokenSaleResponse, w http.ResponseWriter) {
	if sign.VerifyNotifySignData() {
		//verified
		if res.RespCode == "0000" { //创建成功
			//保存token, RegisterUserId 用户id,TokenInvalidDate token有效期, MaskedPan卡号4位
			go service.SaveToken(res.RegisterUserId, res.Token, res.TokenInvalidDate, res.MaskedPan)
			orderId, _ := strconv.Atoi(res.TransId)
			status := api_model.OrderStatus(res.Status)
			o := common.Order{OrderId: orderId, Status: status}
			go o.Update("status", "order_id", o.OrderId)

			w.Write(utils.ApiRetrun(http.StatusOK, "创建成功", nil))
		} else {
			w.Write(utils.ApiRetrun(http.StatusOK, res.RespMsg, nil))
		}
	} else {
		//verify failed
		w.Write(utils.ApiRetrun(http.StatusInternalServerError, "服务器内部错误", nil))
	}
}
