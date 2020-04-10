package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/wechat/routers"
)

func main()  {
	gin.SetMode(gin.DebugMode)
	router := routers.InitPayRouter()
	//err := databases.Orm.Sync2(new(model.Test1))
	//fmt.Println("err:", err)
	//defer databases.Orm.Close()
	router.Run(":6019")

}
