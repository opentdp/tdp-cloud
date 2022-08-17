package master

import (
	"os"

	"github.com/gin-gonic/gin"

	"tdp-cloud/cmd/args"

	"tdp-cloud/internal/api"
	"tdp-cloud/internal/dborm"
	"tdp-cloud/internal/helper"
	"tdp-cloud/internal/migrator"

	"tdp-cloud/front"
)

var engine *gin.Engine

func Create(addr string, mids ...func(*gin.Engine)) {

	// 连接数据库

	dborm.Connect(args.Dsn)

	// 实施自动迁移

	migrator.Start()

	// 启动WEB服务器

	if os.Getenv("TDP_DEBUG") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.Default()

	api.Router(engine)
	front.Router(engine)

	helper.WebServer(addr, engine)

}
