package server

import (
	"log"
	"strings"

	"github.com/kardianos/service"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/httpd"
	"tdp-cloud/module/migrator"
)

type program struct{}

func (p *program) Start(s service.Service) error {

	log.Println("TDP Server start")

	go p.run()
	return nil

}

func (p *program) Stop(s service.Service) error {

	log.Println("TDP Server stop")

	return nil

}

func (p *program) run() {

	// 获取参数
	dsn := viper.GetString("server.dsn")
	listen := viper.GetString("server.listen")

	if !strings.HasPrefix(dsn, "/") {
		dsn = viper.GetString("dataset.dir") + "/" + dsn
	}

	// 连接数据库
	dborm.Connect(dsn)

	// 实施自动迁移
	migrator.Deploy()

	// 启动HTTP服务
	httpd.Start(listen, args.Efs)

}
