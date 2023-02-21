package httpd

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"tdp-cloud/api"
)

func Start(addr string, efs *embed.FS) {

	if viper.GetBool("debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	Server(addr, Engine(efs))

}

func Engine(efs *embed.FS) *gin.Engine {

	// 初始化
	engine := gin.Default()

	// 接口路由
	api.Router(engine)

	// 静态文件路由
	fs, _ := fs.Sub(efs, "front")
	engine.StaticFS("/ui", http.FS(fs))

	// 默认首页路由
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	return engine

}
