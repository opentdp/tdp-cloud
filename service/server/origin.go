package server

import (
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/go-helper/httpd"

	"tdp-cloud/api"
	"tdp-cloud/cmd/args"
	"tdp-cloud/module/certbot"
	"tdp-cloud/module/crontab"
	"tdp-cloud/module/migrator"
)

func origin() {

	dbConnect()

	go certbot.Daemon()
	go crontab.Daemon()

	go httpServer()

}

func dbConnect() {

	// 连接数据库
	dborm.Connect(&dborm.Config{
		Type:     args.Database.Type,
		Host:     args.Database.Host,
		User:     args.Database.User,
		Password: args.Database.Passwd,
		DbName:   args.Database.Name,
		Option:   args.Database.Option,
	})

	// 实施自动迁移
	migrator.Deploy()

}

func httpServer() {

	// 初始化
	engine := httpd.Engine(args.Debug)

	// 接口路由
	api.Router(engine)

	// 上传文件路由
	httpd.Static("/upload", args.Dataset.Dir+"/upload")

	// 前端文件路由
	httpd.StaticEmbed("/", "front", args.Efs)

	// 启动服务
	httpd.Server(args.Server.Listen)

}
