package main

import (
	"tdp-cloud/api"
	"tdp-cloud/front"

	"tdp-cloud/core/cli"
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/serve"
	"tdp-cloud/core/slave"

	"tdp-cloud/core/migrator"
)

func main() {

	cli.Flags()

	if cli.Master == "" {
		server()
	} else {
		client()
	}

}

func server() {

	// 连接数据库
	dborm.Connect(cli.Dsn)

	// 实施自动迁移
	migrator.Start()

	// 创建HTTP服务
	serve.Create(cli.Listen, api.Router, front.Router)

}

func client() {

	// 连接服务端
	slave.Connect(cli.Master)

}
