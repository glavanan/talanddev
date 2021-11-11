package router

import (
	"github.com/gin-gonic/gin"
)

//GetRouter retunr a router with routes initialized
func GetRouter() *gin.Engine {
	router := gin.Default()
	return router
}
