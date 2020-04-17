package api_model

/*
* 使用token消费
 */

type TokenConsumeRequest struct {
	CommonRequest
	ThreeDSType   string `json:"threeDSType"`   // len 3 		交易类型3DS 交易类型，用于区分 3DS1.0 和 2.0 使用 3D 交易必传。示例：1.0,2.0 如果 payMethod 为 judgement，传 2.0
	DeviceChannel string `json:"deviceChannel"` // len 10 		设备的通道类型， 示例：app ,browser
	AccessType    string `json:"accessType"`    // len 16	M	s2s
	//TransChannel string `json:"transChannel"` // len 20 M
	PayMethod      string `json:"payMethod"`      // len 20 O
	TransTimeout   string `json:"transTimeout"`   // len 6	M	订单有效时长，整数格式，单位分钟（min）
	CarrierId      string `json:"carrierId"`      // len 256	M	发生交易的网站域名或者 APP 名称，需在 iPayLinks 备案成功。商户可 MPS 商户后台申请
	Currency       string `json:"currency"`       // len	3	M	交易币种，符合 ISO 4217，需要大写
	TransAmt       string `json:"transAmt"`       // len    12 	M	交易金额，金额需精确到最小金额单位，详见 附录<金额处理与货币代码>
	SettleCurrency string `json:"settleCurrency"` // len 3
	SubAccountData string `json:"subAccountData"` // len 1024
	GoodsName      string `json:"goodsName"`      // len 256 	M	商品名称，多个商品，用“|”隔开
	GoodsInfo      string `json:"goodsInfo"`      // len 2000	O	商品描述，多个商品，用“|”隔开。如果参数 过长，需进行截取

	RegisterUserId string `json:"registerUserId"` // len	32	M	用户在商户端的注册 Id，是用户在商户端的唯一标识
	Token          string `json:"token"`          // len	32	M	用户支付下单时选择卡片对应的 token 值
	Cvv            string `json:"cvv"`            // len (3,4) C 卡片 CVV 验证码，cardVoucher 无值时必填
	CardVoucher    string `json:"cardVoucher"`    // len 64 O	通过 iPayLinks 的 JS SDK 创建的 cardVoucher，	以代替卡片 CVV 验证码
	BillingDesc    string `json:"billingDesc"`    // len 36 O
	//持卡人账单中显示该笔订单的账单名称，格式：
	//"Merchant DBA name" + "*" + "DBA city"。
	//1) Merchant DBA name 账单名称，最长 22 字
	//符；
	//2) *；
	//3) DBA city 可填写商户城市英文名或者联系电话，最长支持 13 字符；如果需要支持动态账单显示，可联系cs@ipaylinks.com
	RiskAssessmentOnly string `json:"riskAssessmentOnly"` // len 5 	O
	RiskInfo           string `json:"riskInfo"`           // len 15000 O 风控信息，json 格式。字段详见<riskInfo>
	Dcc                string `json:"dcc"`                // len 3 O
	RedirectUrl        string `json:"redirectUrl"`        // len 512	商户交易结果回调地址，用于浏览器跳转。需要做页面跳转的支付方式需传送,详见<支付方式列表>
}

type TokenConsumeResponse struct {
	CommonResponse
	CompleteTime   string `json:"completeTime"`   // len 14 O 交易处理完成时间，yyyyMMddHHmmss， eg:20170607125959
	OrderId        string `json:"orderId"`        // len 24 O iPayLinks 订单号
	RiskResult     string `json:"riskResult"`     // lne 1024 O 风控返回信息，json 格式，详见<riskResult>
	ThreeDSResult  string `json:"threeDsResult"`  // len 512 C 3D 交易返回 , 详见<threeDSResult>
	PayMethod      string `json:"payMethod"`      // len 20 C 同请求报文
	Currency       string `json:"currency"`       // len 3 M 同请求报文
	SettleCurrency string `json:"settleCurrency"` // len 3 O 同请求报文
	TransAmt       string `json:"transAmt"`       //len 12 M

	//同步响应才有这个字段
	PayMethodResp string `json:"payMethodResp"` // len 1024 M 支付方式的返回信息，json 格式。详见 <payMethodResp 参数补充说明>
}
