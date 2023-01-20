package server

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/httpd"
	"tdp-cloud/internal/api"
	"tdp-cloud/internal/dborm"
	"tdp-cloud/internal/migrator"
)

var frontFS *embed.FS

func Create(vfs *embed.FS) {

	frontFS = vfs

	// 连接数据库
	dborm.Connect(vDsn)

	// 实施自动迁移
	migrator.Deploy()

	// 设置调试模式
	if os.Getenv("TDP_DEBUG") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 启动HTTP服务器
	httpd.WebServer(vListen, newEngine())

}

func newEngine() *gin.Engine {

	engine := gin.Default()

	api.Router(engine)

	fs, _ := fs.Sub(frontFS, "front")
	engine.StaticFS("/ui", http.FS(fs))

	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	return engine

}
