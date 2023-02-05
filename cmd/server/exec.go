package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/httpd"
	"tdp-cloud/module/migrator"
	"tdp-cloud/module/service"
)

func Execute(cmd *cobra.Command, params []string) {

	switch svc {
	case "install":
		service.Server().Install()
	case "uninstall":
		service.Server().Uninstall()
	case "":
		start()
	}

}

func start() {

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
