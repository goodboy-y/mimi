package main

import (
	"log"
	gt "mimi/pkg"
)

func main() {
	ipt := &gt.Intercept{
		Ip: "0.0.0.0:8111",
		HttpPackageFunc: func(pack *gt.HttpPackage) {
			log.Println("json = ", pack.Json())
			log.Printf("html = %s. gzip = %v", pack.Html(), pack.Gzip)
		},
	}
	// 启动服务
	ipt.RunServer()
}
