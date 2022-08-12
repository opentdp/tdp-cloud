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

	if cli.Agent == "" {
		server()
	} else {
		client()
	}

}

func client() {

	// 连接服务端
	slave.Connect(cli.Agent)

}

func server() {

	// 连接数据库
	dborm.Connect(cli.Dsn)

	// 实施自动迁移
	migrator.Migrate()

	// 创建HTTP服务
	serve.Create(cli.Address, api.Router, front.Router)

}
