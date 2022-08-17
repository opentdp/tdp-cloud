package master

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"tdp-cloud/cmd/args"

	"tdp-cloud/helper/httpd"
	"tdp-cloud/internal/api"
	"tdp-cloud/internal/dborm"
	"tdp-cloud/internal/migrator"
)

var engine *gin.Engine

func Create(vfs *embed.FS) {

	// 连接数据库

	dborm.Connect(args.Dsn)

	// 实施自动迁移

	migrator.Start()

	// 启动HTTP服务器

	if os.Getenv("TDP_DEBUG") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	httpd.WebServer(args.Listen, newEngine(vfs))

}

func newEngine(vfs *embed.FS) *gin.Engine {

	engine = gin.Default()

	api.Router(engine)

	fs, _ := fs.Sub(vfs, "front")
	engine.StaticFS("/ui", http.FS(fs))

	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	return engine

}
