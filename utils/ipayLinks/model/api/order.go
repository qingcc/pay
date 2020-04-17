package api_model

import (
	"github.com/qingcc/yi/commobj"
)

type PayCreateRequest struct {
	CommonRequest
	//Version string `json:"version"`				// len 3 	M	eg. 2.0
	ThreeDSType   string `json:"threeDSType"`   // len 3 		交易类型3DS 交易类型，用于区分 3DS1.0 和 2.0 使用 3D 交易必传。示例：1.0,2.0 如果 payMethod 为 judgement，传 2.0
	DeviceChannel string `json:"deviceChannel"` // len 10 		设备的通道类型， 示例：app ,browser
	//MerchantId string `json:"merchantId"` 		// len 16   M	商户号，由 iPayLinks 分配给商户的唯一 ID
	AccessType string `json:"accessType"` // len 16	M	s2s

	//商户交易流水号，建议每次请求的 transId 唯一， 有利于后续订单查询和对账。对交易状态为failed 订单的 transId，iPayLinks 支持再次提交
	//TransId   string `json:"transId"`//len 64 M
	//TransType string `json:"transType"` // len 18 M sale

	// 支付类型。国际信用卡，传送‘credit_card’。
	//其他渠道枚举值详见附录<支付方式列表>。如
	//果需要开通其他支付方式，可联系cs@ipaylinks.com
	TransChannel string `json:"transChannel"` // len 20 M

	//支付方式，如果需要走该 transChannel 下非默
	//认的 payMethod 时需传送。例：credit_card
	//normal：非 3D 交易
	//3d：3D 交易
	//judgement：intelligent 3d
	//如果不传值，默认走非 3D 交易
	//各 transChannel 下的 payMethod 详见附录<支付
	//方式列表>
	PayMethod    string `json:"payMethod"`    // len 20 O
	TransTimeout string `json:"transTimeout"` // len 6	M	订单有效时长，整数格式，单位分钟（min）
	CarrierId    string `json:"carrierId"`    // len 256	M	发生交易的网站域名或者 APP 名称，需在 iPayLinks 备案成功。商户可 MPS 商户后台申请
	Currency     string `json:"currency"`     // len	3	M	交易币种，符合 ISO 4217，需要大写
	TransAmt     string `json:"transAmt"`     // len    12 	M	交易金额，金额需精确到最小金额单位，详见 附录<金额处理与货币代码>

	//订单结算币种，符合 ISO 4217，需要大写，详
	//见附录<金额处理与货币代码>。默认按照合同
	//结算，如果商户希望动态传送结算币种，可联
	//系 cs@ipaylinks.com
	SettleCurrency string `json:"settleCurrency"` // len 3

	//分账串，各子商户分账金额累计需小于订单的
	//transAmt，json 格式，详见<subAccountData>。
	//如果商户有分账需求，可联系 cs@ipaylinks.com
	SubAccountData string `json:"subAccountData"` // len 1024
	GoodsName      string `json:"goodsName"`      // len 256 	M	商品名称，多个商品，用“|”隔开
	GoodsInfo      string `json:"goodsInfo"`      // len 2000	O	商品描述，多个商品，用“|”隔开。如果参数 过长，需进行截取

	//支付方式的支付要素、支付信息，json 格式。
	//每 种 支 付 方 式 的 payMethodInfo 字 段 详 见
	//<payMethodInfo 参数补充说明>
	PayMethodInfo string `json:"payMethodInfo"`

	//true:只进行风险评估；
	//false:既作风险评估，也进行交易；
	//[Note：默认值为 false]
	RiskAssessmentOnly string `json:"riskAssessmentOnly"` // len 5 	O

	MpiData  string `json:"mpiData"`  // len 600	详见<mpiData 参数补充说明>
	RiskInfo string `json:"riskInfo"` // len 15000 O 风控信息，json 格式。字段详见<riskInfo>
	Dcc      string `json:"dcc"`      // len 3 O

	//如果该笔交易需要走 DCC，传送：dcc。在进行
	//DCC 交易前，需要进行<DCC 锁汇>。DCC 目前
	//仅适用于 credit_card 的交易
	//NotifyUrl string `json:"notifyUrl"`// len 512 O 商户服务器接收 iPayLinks 异步结果的通知地址，可通过 MPS 商户后台配置
	RedirectUrl string `json:"redirectUrl"` // len 512	商户交易结果回调地址，用于浏览器跳转。需要做页面跳转的支付方式需传送,详见<支付方式列表>
	//ReqReserved string `json:"reqReserved"`// len 256 O 商户扩展字段，会原样返回
	//Reserved string `json:"reserved"` // len 2048 O 保留字段
	//Charset string `json:"charset"` // len 5 M utf-8
	//SignType string `json:"signType"` // lne 4 M 报文签名类型：MD5/RSA
	//Sign string `json:"sign"` // len 256 M 报文签名串
}

//下单响应
type PayResponse struct {
	CommonResponse
	CompleteTime   string `json:"completeTime"`   // len 14 O 交易处理完成时间，yyyyMMddHHmmss， eg:20170607125959
	OrderId        string `json:"orderId"`        // len 24 O iPayLinks 订单号
	RiskResult     string `json:"riskResult"`     // lne 1024 O 风控返回信息，json 格式，详见<riskResult>
	ThreeDSResult  string `json:"threeDsResult"`  // len 512 C 3D 交易返回 , 详见<threeDSResult>
	TransChannel   string `json:"transChannel"`   // len 20 M 同请求报文
	PayMethod      string `json:"payMethod"`      // len 20 C 同请求报文
	Currency       string `json:"currency"`       // len 3 M 同请求报文
	SettleCurrency string `json:"settleCurrency"` // len 3 O 同请求报文

	//同步响应才有这个字段
	PayMethodResp string `json:"payMethodResp"` // len 1024 M 支付方式的返回信息，json 格式。详见 <payMethodResp 参数补充说明>
}

func OrderStatus(status string) (orderStatus commobj.OrderStatus) {
	switch status {
	case "success":
		orderStatus = commobj.ORDER_STATUS_CONFIRM
	case "failed":
		orderStatus = commobj.ORDER_STATUS_FAILED_RESUBMIT
	case "expired":
		orderStatus = commobj.ORDER_STATUS_FAILED_RESUBMIT
	case "canceled":
		orderStatus = commobj.ORDER_STATUS_CANCELED
	case "received", "pending":
		orderStatus = commobj.ORDER_STATUS_APPENDING
	default: //pending_review或未识别的状态，默认需要人工检查
		orderStatus = commobj.ORDER_STATUS_CHECK_NEEDED
	}
	return
}

/*
* 订单状态status
* success 			Y	交易成功--预授权可做取消 按交易成功处理，不可重新提交
* failed 			Y	交易失败 按交易失败处理，可重新提交
* expired 			Y	订单过期 按订单过期处理，可重新提交
* canceled 			Y	交易取消 按交易成功处理(交易成功，但已被撤销)，不可重新提交
* received 			N	初始状态 按交易中处理，不可重新提交
* pending 			N	交易中 按交易中处理，不可重复提交
* pending_review	N	等待审核 交易提交成功，需商户人工审核，不可重复提交
 */
