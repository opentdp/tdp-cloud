package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"

	"tdp-cloud/api/cam"
	"tdp-cloud/api/dnspod"
	"tdp-cloud/api/lighthouse"
	"tdp-cloud/api/monitor"

	"tdp-cloud/api/user"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.JSON())

	{
		// cloud api

		cloud := api.Group("/cloud")

		{
			cam.Router(cloud)
			dnspod.Router(cloud)
			lighthouse.Router(cloud)
			monitor.Router(cloud)
		}

		// local api

		local := api.Group("/local")

		{
			user.Router(local)
		}
	}

}
