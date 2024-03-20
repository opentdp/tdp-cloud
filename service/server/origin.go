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
		Type:     args.Gormio.Type,
		Host:     args.Gormio.Host,
		User:     args.Gormio.User,
		Password: args.Gormio.Passwd,
		DbName:   args.Gormio.Name,
		Option:   args.Gormio.Option,
	})

	// 设置默认参数
	if args.Gormio.Type == "sqlite" {
		dborm.Db.Exec("PRAGMA foreign_keys=ON;")
		dborm.Db.Exec("PRAGMA journal_mode=WAL;")
		dborm.Db.Exec("PRAGMA busy_timeout=5000;")
	}

	// 实施自动迁移
	migrator.Deploy()

}

func httpServer() {

	// 初始化
	engine := httpd.Engine(args.Debug)

	// 接口路由
	api.Router(engine)

	// 上传文件路由
	httpd.Static("/upload", args.Assets.Dir+"/upload")

	// 前端文件路由
	httpd.StaticEmbed("/", "front", args.Efs)

	// 启动服务
	httpd.Server(args.Server.Listen)

}
