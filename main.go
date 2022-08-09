package main

import (
	"tdp-cloud/api"
	"tdp-cloud/front"

	"tdp-cloud/core/cli"
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/serve"

	"tdp-cloud/core/client"
)

func main() {

	cli.Flags()

	if cli.Agent == "" {
		server()
	} else {
		agent()
	}

}

func agent() {

	// 连接服务端
	client.Connect(cli.Agent)

}

func server() {

	// 连接数据库
	dborm.Connect(cli.Dsn)

	// 创建HTTP服务
	serve.Create(cli.Address, api.Router, front.Router)

}
