package utils

import (
	"encoding/json"
	"github.com/qingcc/wechat/http/model"
)

func ApiRetrun(code int, message string, data interface{}) []byte {
	b, _ := json.Marshal(model.ApiCommonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
	return b
}
