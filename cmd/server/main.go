package server

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"tdp-cloud/api"
	"tdp-cloud/cmd"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/httpd"
	"tdp-cloud/module/migrator"
)

func Create() {

	// 连接数据库
	dborm.Connect(vDsn)

	// 实施自动迁移
	migrator.Deploy()

	// 设置调试模式
	if os.Getenv("TDP_DEBUG") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 启动HTTP服务
	httpd.WebServer(vListen, newEngine())

}

func newEngine() *gin.Engine {

	engine := gin.Default()

	// 接口路由
	api.Router(engine)

	// 静态文件路由
	fs, _ := fs.Sub(cmd.FrontFS, "front")
	engine.StaticFS("/ui", http.FS(fs))

	// 默认首页路由
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	return engine

}
