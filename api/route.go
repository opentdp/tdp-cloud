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
		// cloud api

		cloud := api.Group("/cloud")

		cloud.Use(midware.Auth())
		cloud.Use(midware.Secret())

		{
			cam.Router(cloud)
			dnspod.Router(cloud)
			lighthouse.Router(cloud)
		}

		// local api

		local := api.Group("/local")

		{
			user.Router(local)
		}
	}

}
