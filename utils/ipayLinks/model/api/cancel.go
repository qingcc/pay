package api_model

type CancelOrder struct {
	CommonRequest
	TransAmt       string `json:"transAmt"`       // len 12 M 交易金额，金额需精确到最小金额单位，详见 附录<金额处理与货币代码>
	OrigTransId    string `json:"origTransId"`    //len	64	M	商户原始交易流水号，即关联 sale/capture 交易的 transId
	SubAccountData string `json:"subAccountData"` // len 1024
}

type CancelResponse struct {
	CommonResponse
	OrigTransId  string `json:"origTransId"`  //len	64	M	商户原始交易流水号，即关联 sale/capture 交易的 transId
	TransAmt     string `json:"transAmt"`     // len 12 M 交易金额，金额需精确到最小金额单位，详见 附录<金额处理与货币代码>
	Currency     string `json:"currency"`     // len	3	M	交易币种，符合 ISO 4217，需要大写
	CompleteTime string `json:"completeTime"` // len 14 O 交易处理完成时间，yyyyMMddHHmmss， eg:20170607125959
	OrderId      string `json:"orderId"`      // len 24 O iPayLinks 订单号
}
