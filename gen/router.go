package gen

import (
	"io/ioutil"
	"strings"
)

func GenRouter(module string) error {
	code := tplRouter
	code = strings.Replace(code, "<module>", module, -1)

	file := "router/router.go"
	if err := ioutil.WriteFile(file, []byte(code), 0644); err != nil {
		return err
	}
	return nil
}

var tplRouter = `package router

import (
	"github.com/gin-gonic/gin"
	"<module>/handler"
	"<module>/middleware"
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
`
