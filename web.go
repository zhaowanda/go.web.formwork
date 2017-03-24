package web

import (
	fast "github.com/valyala/fasthttp"
	"log"
	"os"
	"github.com/zhaowanda/go.web.formwork/web/router"
	"github.com/zhaowanda/go.web.formwork/yaml"
        "github.com/zhaowanda/go.web.formwork/config"
)

//  binlog 同步web配置端
//  提供restful接口,页面后期进行整合

func WebMain() {
	ipAddr, err := yaml.GetConfig("server.boot.ListenAndServe", config.YamlConfig)
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
	log.Fatal(fast.ListenAndServe(ipAddr.(string), webrouter.InitRouter().Handler))
}
