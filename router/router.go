package router

import (
	"github.com/gin-gonic/gin"
	"github.com/glavanan/talanddev/handler"
)

//GetRouter retunr a router with routes initialized
func GetRouter() *gin.Engine {
	router := gin.Default()
	router.POST("convert", handler.Convert)
	return router
}
