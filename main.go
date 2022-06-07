package main

import (
	"flag"

	"tdp-cloud/api"
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/serve"
	"tdp-cloud/front"
)

var (
	dsn    string
	listen string
)

func main() {

	flags()

	// 连接数据库

	dborm.Connect(dsn)

	// 创建HTTP服务

	engine := serve.Create()

	api.Router(engine)
	front.Router(engine)

	serve.Listen(listen)

}

func flags() {

	flag.StringVar(&dsn, "dsn", "cloud.db", "数据库存储路径")

	flag.StringVar(&listen, "listen", ":7800", "服务器监听地址和端口")

	flag.Parse()
}
