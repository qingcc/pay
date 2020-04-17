package model

import "github.com/qingcc/yi/commobj"

type ApiOrderResponse struct {
	OrderStatus commobj.OrderStatus `json:"order_status"`
}
