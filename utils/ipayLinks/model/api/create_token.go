package api_model

type CreateTokenRequest struct {
	CommonRequest
	ThreeDSType    string `json:"threeDSType"`    // len 3 		交易类型3DS 交易类型，用于区分 3DS1.0 和 2.0 使用 3D 交易必传。示例：1.0,2.0 如果 payMethod 为 judgement，传 2.0
	DeviceChannel  string `json:"deviceChannel"`  // len 10 		设备的通道类型， 示例：app ,browser
	AccessType     string `json:"accessType"`     // len 16	M	s2s
	PayMethod      string `json:"payMethod"`      // len 20 O
	TransTimeout   string `json:"transTimeout"`   // len 6	M	订单有效时长，整数格式，单位分钟（min）
	CarrierId      string `json:"carrierId"`      // len 256	M	发生交易的网站域名或者 APP 名称，需在 iPayLinks 备案成功。商户可 MPS 商户后台申请
	PayMethodInfo  string `json:"payMethodInfo"`  //len 600 M	credit_card 的支付要素、支付信息，json 格式。详见<credit_card 的 payMethodInfo 参数补充说明>
	RegisterUserId string `json:"registerUserId"` //len	32	M	用户在商户端的注册 Id，是用户在商户端的唯一标识
	RiskInfo       string `json:"riskInfo"`       // len	15000	O	风控信息，json 格式。字段详见<riskInfo>
	RedirectUrl    string `json:"redirectUrl"`    // len 512	商户交易结果回调地址，用于浏览器跳转。需要做页面跳转的支付方式需传送,详见<支付方式列表>
}

type CreateTokenResponse struct {
	CommonResponse
	PayMethod        string `json:"payMethod"`        // len 20 O
	CompleteTime     string `json:"completeTime"`     // len 14 O 交易处理完成时间，yyyyMMddHHmmss， eg:20170607125959
	OrderId          string `json:"orderId"`          // len 24 O iPayLinks 订单号
	MaskedPan        string `json:"maskedPan"`        //len	13	C	卡号后 4 位
	CardScheme       string `json:"cardScheme"`       //len	16	C	卡组织，值： visa/mastercard/jcb/american_express/diners_cl 		ub/discover
	RegisterUserId   string `json:"registerUserId"`   //len	32	M	同请求报文
	Token            string `json:"token"`            //len	32	C	该 registerUserId 下该卡卡信息的 token 值，可 代替卡信息用于后续交易。交易成功才返回
	TokenInvalidDate string `json:"tokenInvalidDate"` //len	8	C	token 的有效期，交易成功才返回。 yyyyMMdd，eg:20170607
	RiskResult       string `json:"riskResult"`       //len	1024	O	风控返回信息，json 格式，详见<riskResult>
	ThreeDSResult    string `json:"threeDSResult"`    //len	512	C	3D 交易返回 , 详见<threeDSResult>

	//同步响应才有这个字段
	PayMethodResp string `json:"payMethodResp"` // len 1024 M 支付方式的返回信息，json 格式。详见 <payMethodResp 参数补充说明>
}
