package main

import (
	"tdp-cloud/api"
	"tdp-cloud/front"

	"tdp-cloud/core/cli"
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/serve"

	"github.com/gin-gonic/gin"
)

func main() {

	cli.Flags()

	// 连接数据库

	dborm.Connect(cli.Dsn)

	// 创建HTTP服务

	serve.Create(cli.Address, func(engine *gin.Engine) {
		api.Router(engine)
		front.Router(engine)
	})

}
