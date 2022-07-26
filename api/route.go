package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"

	"tdp-cloud/api/member"
	"tdp-cloud/api/tat"

	"tdp-cloud/api/cam"
	"tdp-cloud/api/dnspod"
	"tdp-cloud/api/lighthouse"
	"tdp-cloud/api/monitor"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.JSON())

	{
		// local api

		local := api.Group("/local")

		{
			member.Router(local)
			tat.Router(local)
		}

		// cloud api

		cloud := api.Group("/cloud")

		{
			cam.Router(cloud)
			dnspod.Router(cloud)
			lighthouse.Router(cloud)
			monitor.Router(cloud)
		}
	}

}
