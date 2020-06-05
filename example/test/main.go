package main

import (
	"github.com/gin-gonic/gin"
	"test/config"
	"test/router"
)

func main() {
	// config
	config.Init()

	// mode
	gin.SetMode(config.App.Mode)

	// router
	r := router.Init()

	// listen and serve
	r.Run(config.App.ServiceAddr) 
}