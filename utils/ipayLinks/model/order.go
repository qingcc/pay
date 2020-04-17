package common

import (
	"fmt"
	"github.com/qingcc/yi/commobj"
	"github.com/qingcc/yi/database/mysqldb"
	"log"
)

type Order struct {
	//todo 订单表具体字段
	OrderId int                 `json:"order_id"`
	Status  commobj.OrderStatus `json:"status"`
}

var errorRetryTimes = 3 //更新操作执行失败时，重试次数

func (o *Order) Update(updateFields string, whereConn string, val ...interface{}) (ok bool) {
	sess := mysqldb.GetRDBConn()
	errorTime := 0
	for {
		has, err := sess.Table(Order{}).Where(whereConn, val).Update(updateFields)
		if err != nil {
			log.Println(fmt.Sprintf("update table failed, err:%s, %s", err.Error(), whereConn), val)
			if errorTime < errorRetryTimes {
				errorTime++
				continue //update sql执行失败，重新执行
			}
		}
		if has > 0 {
			ok = true
		}
		break
	}
	return
}
