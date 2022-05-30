package main

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/api"
	"tdp-cloud/core/serve"
	"tdp-cloud/front"
)

func main() {

	engine := gin.Default()

	api.Router(engine)
	front.Router(engine)

	serve.Init(engine)

}
