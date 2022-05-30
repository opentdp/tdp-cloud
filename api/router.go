package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/api/cam"
	"tdp-cloud/api/dnspod"
	"tdp-cloud/api/lighthouse"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	{
		cam.Router(api)
		dnspod.Router(api)
		lighthouse.Router(api)
	}

}
