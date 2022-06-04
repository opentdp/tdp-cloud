package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/api/cam"
	"tdp-cloud/api/dnspod"
	"tdp-cloud/api/lighthouse"

	"tdp-cloud/api/user"

	"tdp-cloud/core/midware"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.JSON())

	{
		cam.Router(api)
		dnspod.Router(api)
		lighthouse.Router(api)

		user.Router(api)
	}

}
