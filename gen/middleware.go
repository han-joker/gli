package gen

import (
	"io/ioutil"
	"strings"
)

func GenMiddleWareCors(module string) error {
	code := tplMiddlewareCors
	code = strings.Replace(code, "<module>", module, -1)

	file := "middleware/cors.go"
	if err := ioutil.WriteFile(file, []byte(code), 0644); err != nil {
		return err
	}
	return nil
}

var tplMiddlewareCors = `package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GenCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("authorization")
	return cors.New(config)
}
`
