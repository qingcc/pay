package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/wechat/utils/ipayLinks/api"
)

func Pay(c gin.Context)  {
	payMethod := c.PostForm("payMethod")
	switch payMethod {
	case "iPayLinks":
		api.PaySync(c.Writer, c.Request)
	}
}
