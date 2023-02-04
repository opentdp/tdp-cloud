package server

import (
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/httpd"
	"tdp-cloud/module/migrator"
)

func Execute() {

	// 获取参数
	dsn := viper.GetString("server.dsn")
	listen := viper.GetString("server.listen")

	// 连接数据库
	dborm.Connect(dsn)

	// 实施自动迁移
	migrator.Deploy()

	// 启动HTTP服务
	httpd.Start(listen, args.EmbedFs)

}
