package gen

import (
	"io/ioutil"
	"strings"
)

func GenEntrance(module string) error {
	code := tplEntrance
	code = strings.Replace(code, "<module>", module, -1)

	//node, err := parser.ParseExpr(code)
	//if err != nil {
	//	return err
	//}
	//fileset := token.NewFileSet()
	//var buf bytes.Buffer
	//if err := format.Node(&buf, fileset, node); err != nil {
	//	return err
	//}

	file := "main.go"
	if err := ioutil.WriteFile(file, []byte(code), 0644); err != nil {
		return err
	}
	return nil
}

var tplEntrance = `package main

import (
	"github.com/gin-gonic/gin"
	"<module>/config"
	"<module>/router"
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
}`
