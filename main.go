package main

import (
	"tdp-cloud/api"
	"tdp-cloud/core/serve"
	"tdp-cloud/front"
)

func main() {

	engine := serve.Create()

	api.Router(engine)
	front.Router(engine)

	serve.Listen(":8000")

}
