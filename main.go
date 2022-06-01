package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"tdp-cloud/api"
	"tdp-cloud/core/serve"
	"tdp-cloud/front"
)

func main() {

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	api.Router(engine)
	front.Router(engine)

	serve.Init(engine, ":8080")

}
