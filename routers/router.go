package routers

import "github.com/gin-gonic/gin"

func InitPayRouter() *gin.Engine {
	router := gin.Default()
	router.Use()
	v1 := router.Group("/pay")
	v1.POST("/test", )
	return
}
