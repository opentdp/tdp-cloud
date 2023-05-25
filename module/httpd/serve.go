package httpd

import (
	"github.com/open-tdp/go-helper/httpd"

	"tdp-cloud/api"
	"tdp-cloud/cmd/args"
)

func Daemon() {

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
