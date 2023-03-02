package httpd

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"tdp-cloud/api"
	"tdp-cloud/cmd/args"
	"tdp-cloud/module/midware"
)

func Start() {

	if args.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	Server(args.Server.Listen, Engine())

}

func Engine() *gin.Engine {

	// 初始化
	engine := gin.New()
	engine.Use(midware.Logger())
	engine.Use(midware.Recovery(true))

	// 接口路由
	api.Router(engine)

	// 静态文件路由
	fs, _ := fs.Sub(args.Efs, "front")
	engine.StaticFS("/ui", http.FS(fs))

	// 默认首页路由
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	return engine

}
