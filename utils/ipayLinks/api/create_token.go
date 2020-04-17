package api

/*
* 只是创建建绑定卡信息的token
 */

import (
	"encoding/json"
	"github.com/qingcc/wechat/utils"
	common "github.com/qingcc/wechat/utils/ipayLinks/model"
	api_model "github.com/qingcc/wechat/utils/ipayLinks/model/api"
	"github.com/qingcc/wechat/utils/ipayLinks/service"
	"github.com/qingcc/wechat/utils/ipayLinks/tools"
	"io/ioutil"
	"net/http"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	rb, _ := ioutil.ReadAll(r.Body)
	req := new(api_model.CreateTokenRequest)
	json.Unmarshal(rb, &req)
	req.Init()

	res := new(api_model.CreateTokenResponse)
	s := common.DefaultNewSign(req)
	tools.IpayLinksSendData(s, res, true)
	sign := &common.Sign{
		Sign: res.Sign,
		Data: res,
	}
	dealCreateTokenResponse(sign, res, w)
}

//异步响应
func CreateTokenSync(w http.ResponseWriter, r *http.Request) {
	rb, _ := ioutil.ReadAll(r.Body)
	cancelReq := new(api_model.CreateTokenResponse)
	json.Unmarshal(rb, &cancelReq)
	sign := &common.Sign{
		Sign: cancelReq.Sign,
		Data: cancelReq,
	}
	dealCreateTokenResponse(sign, cancelReq, w)
}

//处理创建token的返回数据
func dealCreateTokenResponse(sign *common.Sign, res *api_model.CreateTokenResponse, w http.ResponseWriter) {
	if sign.VerifyNotifySignData() {
		//verified
		if res.RespCode == "0000" { //创建成功
			//保存token, RegisterUserId 用户id,TokenInvalidDate token有效期, MaskedPan卡号4位
			go service.SaveToken(res.RegisterUserId, res.Token, res.TokenInvalidDate, res.MaskedPan)
			w.Write(utils.ApiRetrun(http.StatusOK, "创建成功", nil))
		} else {
			w.Write(utils.ApiRetrun(http.StatusOK, res.RespMsg, nil))
		}
	} else {
		//verify failed
		w.Write(utils.ApiRetrun(http.StatusInternalServerError, "服务器内部错误", nil))
	}
}
