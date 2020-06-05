package gen

import (
	"fmt"
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

func GenRouterResource(resource string) error {
	code := tplRouterResource
	code = strings.Replace(code, "<resource>", fmtUri(resource), -1)
	code = strings.Replace(code, "<method>", fmtMethod(resource), -1)
	fmt.Println(code)

	return nil
}

func fmtUri(str string) string {
	s := strings.Replace(strings.ToLower(str), "_", "-", -1)
	return s
}
func fmtMethod(str string) string {
	ss := []string{str}
	if strings.Contains(str, "-") {
		ss = strings.Split(strings.ToLower(str), "-")
	} else if strings.Contains(str, "_") {
		ss = strings.Split(strings.ToLower(str), "_")
	}
	for i, _ := range ss {
		ss[i] = UpperFirst(ss[i])
	}
	return strings.Join(ss, "")
}

func UpperFirst(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
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
var tplRouterResource = `
r.GET("/<resource>", handler.<method>List)
r.GET("/<resource>/:id", handler.<method>Info)
r.POST("/<resource>", handler.<method>Create)
r.PUT("/<resource>/:id", handler.<method>Edit)
r.DELETE("/<resource>/:id", handler.<method>Remove)
`
