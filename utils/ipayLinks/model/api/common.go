package api_model

import common "github.com/qingcc/wechat/utils/ipayLinks/model"

type Common struct {
	Version    string `json:"version"`    // len 3 M 2.0
	MerchantId string `json:"merchantId"` // len 16 M 商户号，由 iPayLinks 分配给商户的唯一 ID

	//商户交易流水号，建议每次请求的 transId 唯一， 有利于后续订单查询和对账。对交易状态为failed 订单的 transId，iPayLinks 支持再次提交
	TransId     string `json:"transId"`     // len 64 M
	TransType   string `json:"transType"`   // len 18 M Sale
	ReqReserved string `json:"reqReserved"` // len 1024 O 商户扩展字段，会原样返回
	Reserved    string `json:"reserved"`    // len 2048 O 保留字段
}

type CommonRequest struct {
	Common
	Charset   string `json:"charset"`   // len 5 M utf-8
	NotifyUrl string `json:"notifyUrl"` // len 512 O 商户服务器接收 iPayLinks 异步结果的通知地址，可通过 MPS 商户后台配置
}

type CommonResponse struct {
	SignType string `json:"signType"` // len 4 M 报文签名类型：MD5/RSA
	Sign     string `json:"sign"`     // len 256 M 报文签名串
	RespCode string `json:"respCode"` // len 4 M 返回码，详见<交易返回码>
	RespMsg  string `json:"respMsg"`  // len 256 M 返回码描述，详见<交易返回码>
	Status   string `json:"status"`   // len 20 M 交易状态，订单的交易结果以订单交易状态为 准。具体规则详见<交易状态说明>
	Common
}

func (o *CommonRequest) Init() {
	o.Version = "2.0"
	o.MerchantId = common.MerchantId
	o.NotifyUrl = common.NotifyUrl
	o.Charset = "utf-8"
}

/*
* RespCode 响应状态码
*
* 返回码	返回消息（英文）		客户端显示(建议)			备注
* 0000 	Approved or completed	successfully		承兑或交易处理成功；
* 0300 		Request accepted 	The transaction was	请求申请成功，只表示系统接受结果，不代表最终系统处理结果；
* 						submitted successfully,wait to confirm
* 0503 		Revoke succeed 							交易撤销操作成功；
* 3118 		Partial approval 						交易成功，部分金额批准；(未开通，不会出现)
* 3200 Transaction is pending 	The transaction was	交易待定，需要商户参与处理：	1.动态预授权成功，需等待商户审核；
* 							submitted successfully,wait to confirm.				2.S2S 交互，动态 3D 需要商户重定向发卡行页面；
* 																				3.收到银行拒付交易，等待商户处理；
* 0055 Revoke completed 							预授权撤销时原预授权申请订单状态更新为撤销完成
 */
