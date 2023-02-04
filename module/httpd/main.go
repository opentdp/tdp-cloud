package httpd

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"tdp-cloud/api"
	"tdp-cloud/cmd/args"
)

func Start(addr string) {

	if os.Getenv("TDP_DEBUG") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	Server(addr, Engine())

}

func Engine() *gin.Engine {

	// 初始化
	engine := gin.Default()

	// 接口路由
	api.Router(engine)

	// 静态文件路由
	fs, _ := fs.Sub(args.FrontFS, "front")
	engine.StaticFS("/ui", http.FS(fs))

	// 默认首页路由
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	return engine

}
