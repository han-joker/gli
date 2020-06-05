package router

import (
	"github.com/gin-gonic/gin"
	"test/handler"
	"test/middleware"
)

func Init() *gin.Engine {
	r := gin.Default()
	
	// cors
	r.Use(middleware.GenCors())
	
	// ping
	r.GET("/ping", handler.Ping)
	
	// return
	return r
}
