package main

import (
	"tdp-cloud/api"
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/serve"
	"tdp-cloud/front"
)

func main() {

	// 连接数据库

	dborm.Connect()

	// 创建HTTP服务

	engine := serve.Create()

	api.Router(engine)
	front.Router(engine)

	serve.Listen(":7800")

}
