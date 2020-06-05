package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
)

var App *AppConf

// 配置类型
type AppConf struct {
	AppName     string `yaml:"app_name" default:"A gli app"`
	ServiceAddr string `yaml:"service_addr" default:":8080"`
	SecretKey   string `yaml:"secret_key" default:""`
	Mode   		string `yaml:"mode" default:"debug"`
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
