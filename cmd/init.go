package cmd

import (
	"fmt"
	"github.com/han-joker/gli/gen"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Init struct {
	baseCMD
}

func (c *Init) Run() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("gli init: dir error.")
		return
	}
	fs, err := ioutil.ReadDir(path)
	if err!=nil {
		fmt.Println("gli init: dir error.")
		return
	}
	if len(fs) > 0 {
		fmt.Println("gli init: dir not empty.")
		return
	}

	c.FlagSet.Parse(os.Args[2:])

	if c.FlagSet.NArg() > 1 {
		fmt.Println("gli init: too many arguments")
		return
	}

	module := ""
	if c.FlagSet.NArg() == 1 {
		 module = c.FlagSet.Arg(0)
	} else if c.FlagSet.NArg() == 0 {
		// 基于目录名推测 module
		module = filepath.Base(path)
	}

	// 生成 go.mod
	runtime.Version()
	if err := gen.GenGoMod(module, "1.14"); err != nil {
		log.Println(err)
	} else {
		log.Println("go.mod was generated.")
	}

	// 生成入口文件
	if err := gen.GenEntrance(module); err != nil {
		log.Println(err)
	} else {
		log.Println("main.go was generated.")
	}

	// 生成目录结构
	// router
	if err := gen.GenRouterDir(); err != nil {
		log.Println(err)
	} else {
		log.Println("router dir was generated.")
	}
	// handler
	if err := gen.GenHandlerDir(); err != nil {
		log.Println(err)
	} else {
		log.Println("handler dir was generated.")
	}
	// middleware
	if err := gen.GenMiddlewareDir(); err != nil {
		log.Println(err)
	} else {
		log.Println("middleware dir was generated.")
	}
	// config
	if err := gen.GenConfigDir(); err != nil {
		log.Println(err)
	} else {
		log.Println("config dir was generated.")
	}
	// config
	if err := gen.GenModelDir(); err != nil {
		log.Println(err)
	} else {
		log.Println("model dir was generated.")
	}
	// 生成路由文件 router/router.go
	if err := gen.GenRouter(module); err != nil {
		log.Println(err)
	} else {
		log.Println("router/router.go was generated.")
	}
	// 生成 cors 中间件 middleware/cors.go
	if err := gen.GenMiddleWareCors(module); err != nil {
		log.Println(err)
	} else {
		log.Println("middleware/cors.go was generated.")
	}
	// 生成 handler/ping.go
	if err := gen.GenHandlerPing(); err != nil {
		log.Println(err)
	} else {
		log.Println("handler/ping.go was generated.")
	}
	//生成 config/app.go
	if err := gen.GenConfigApp(); err != nil {
		log.Println(err)
	} else {
		log.Println("config/app.go was generated.")
	}
}

func (c *Init) Usage() {
	fmt.Println("gli init provides initialize the application")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("\t%s\n", "gli init module-name")
}

