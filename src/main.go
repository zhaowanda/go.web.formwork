package main

import (
	"web"
	"web/router"
	"web/handler"
	"log"
	"os"
	"config"
)


func main() {
	//canalgo.Main()
	if config.Err != nil {
		log.Println(config.Err)
		os.Exit(3)
	}
	web.WebMain()
}


// 初始化业务逻辑中的 url 及 handler
func init() {
	webrouter.UrlCache["/hello"] = requesthandler.Hello
}