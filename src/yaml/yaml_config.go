package yaml

import (
	"errors"
	yaml "github.com/wendal/goyaml2"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

var(
	ConfigCache = make(map[string]interface{})
)

// 定义配置文件帮助类
type ConfigHelper struct {
	Config map[string]interface{}
}

// 解析 yaml 配置文件
func InitYamlConfig(path string) (*ConfigHelper, error) {
	config := &ConfigHelper{}
	config.Config = make(map[string]interface{})
	fi, err := os.Open(path)
	defer fi.Close()
	if err != nil {
		log.Println("读取文件出错, error:", err)
		return nil, err
	}
	reader := io.Reader(fi)
	obj, err := yaml.Read(reader)
	if err != nil {
		log.Println("解析yaml文件出错,error:", err)
		return nil, err
	}
	config.Config = obj.(map[string]interface{})
	return config, err
}

// 获取 yaml 配置
func GetConfig(conf string, config *ConfigHelper) (interface{}, error) {
	if len(conf) == 0 {
		return nil, errors.New("conf.name is null")
	}
	result, ok := ConfigCache[conf]
	if ok {
		return result, nil
	}
	if config.Config == nil {
		return nil, errors.New("conf is null")
	}
	obj := config.Config
	arr := strings.Split(conf, ".")
	for _, value := range arr {
		if strings.EqualFold(reflect.Map.String(), reflect.TypeOf(obj).Kind().String()) {
			object := obj[value]
			if strings.EqualFold(reflect.Map.String(), reflect.TypeOf(object).Kind().String()) {
				obj = object.(map[string]interface{})
			} else {
				result = object
			}
		} else {
			result = obj
		}
	}
	return result, nil
}
