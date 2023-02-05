package server

import (
	"log"

	"github.com/kardianos/service"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/httpd"
	"tdp-cloud/module/migrator"
)

type origin struct{}

func (p *origin) Start(s service.Service) error {

	log.Println("Server service start")

	go p.run()
	return nil

}

func (p *origin) Stop(s service.Service) error {

	log.Println("Server service stop")

	return nil

}

func (p *origin) run() {

	// 获取参数
	dsn := viper.GetString("server.dsn")
	listen := viper.GetString("server.listen")

	// 连接数据库
	dborm.Connect(dsn)

	// 实施自动迁移
	migrator.Deploy()

	// 启动HTTP服务
	httpd.Start(listen, args.Efs)

}
