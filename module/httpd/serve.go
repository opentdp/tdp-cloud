package httpd

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-tdp/go-helper/httpd"

	"tdp-cloud/api"
	"tdp-cloud/cmd/args"
)

func Daemon() {

	// 初始化
	engine := httpd.Engine(args.Debug)

	// 接口路由
	api.Router(engine)

	// 前端文件路由
	ui, _ := fs.Sub(args.Efs, "front")
	engine.StaticFS("/ui", http.FS(ui))

	// 上传文件路由
	engine.Static("/upload", args.Dataset.Dir+"/upload")

	// 默认首页路由
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	httpd.Server(args.Server.Listen, engine)

}
