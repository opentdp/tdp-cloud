package server

import (
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/httpd"
	"tdp-cloud/module/migrator"
)

func Create() {

	// 连接数据库
	dborm.Connect(vDsn)

	// 实施自动迁移
	migrator.Deploy()

	// 启动HTTP服务
	httpd.Start(vListen)

}
