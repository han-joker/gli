package gen

import (
	"io/ioutil"
	"strings"
)

func GenConfigApp() error {
	code := tplConfigApp
	code = strings.Replace(code, "<back-qoute>", "`", -1)

	file := "config/app.go"
	if err := ioutil.WriteFile(file, []byte(code), 0644); err != nil {
		return err
	}
	return nil
}

var tplConfigApp = `package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
)

var App *AppConf

// 配置类型
type AppConf struct {
	AppName     string <back-qoute>yaml:"app_name" default:"A gli app"<back-qoute>
	ServiceAddr string <back-qoute>yaml:"service_addr" default:":8080"<back-qoute>
	SecretKey   string <back-qoute>yaml:"secret_key" default:""<back-qoute>
	Mode   		string <back-qoute>yaml:"mode" default:"debug"<back-qoute>
}

func Init() error {
	App = &AppConf{}

	// default
	rftc := reflect.TypeOf(*App)
	rfvc := reflect.ValueOf(App)
	for i := 0; i < rftc.NumField(); i++ {
		tf := rftc.Field(i)
		vf := rfvc.Elem().Field(i)
		vf.SetString(tf.Tag.Get("default"))
	}

	// configuration file
	confFile := "conf/app.yml"
	if _, err := os.Stat(confFile); err != nil {
		return err
	}
	content, err := ioutil.ReadFile(confFile)
	if err != nil {
		return err
	}
	// 使用 JSON 解析 unmarshal
	if err := yaml.Unmarshal(content, App); err != nil {
		// 解析配置文件错误
		return err
	}

	return nil
}
`
