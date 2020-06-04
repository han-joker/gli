package gen

import (
	"io/ioutil"
)

func GenHandlerPing() error {
	code := tplHandlerPing

	file := "handler/ping.go"
	if err := ioutil.WriteFile(file, []byte(code), 0644); err != nil {
		return err
	}
	return nil
}

var tplHandlerPing = `package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "pong",
		"error": nil,
	})
}`
