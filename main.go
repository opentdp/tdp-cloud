package main

import (
	"tdp-cloud/api"
	"tdp-cloud/front"

	"tdp-cloud/core/cli"
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/serve"
)

func main() {

	cli.Flags()

	// 连接数据库

	dborm.Connect(cli.Dsn)

	// 创建HTTP服务

	serve.Create(cli.Address, api.Router, front.Router)

}
