package api_model

type CommonRequest struct {
	Version string `json:"version"`// len 3 M 2.0
	MerchantId string `json:"merchantId"` // len 16 M 同请求报文
	TransId string `json:"transId"` // len 64 M 同请求报文
	TransType string `json:"transType"` // len 18 M Sale
	TransChannel string `json:"transChannel"` // len 20 M 同请求报文
	PayMethod string `json:"payMethod"` // len 20 C 同请求报文
	Currency string `json:"currency"` // len 3 M 同请求报文
	TransAmt string `json:"transAmt"` // len 12 M 同请求报文
	ReqReserved string `json:"reqReserved"` // len 1024 O 同请求报文
	Reserved string `json:"reserved"` // len 2048 O 保留字段
	SettleCurrency string `json:"settleCurrency"` // len 3 O 结算币种，符合 ISO 4217，详见附录<金额处理 与货币代码>
}

type CommonResponse struct {
	SignType string `json:"signType"` // len 4 M 报文签名类型：MD5/RSA
	Sign string `json:"sign"` // len 256 M 报文签名串
	CommonRequest
}