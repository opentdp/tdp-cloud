package cloud

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/api/cloud/cam"
	"tdp-cloud/api/cloud/dnspod"
	"tdp-cloud/api/cloud/lighthouse"
	"tdp-cloud/api/cloud/monitor"
)

func Router(api *gin.RouterGroup) {

	cloud := api.Group("/cloud")

	{
		cam.Router(cloud)
		dnspod.Router(cloud)
		lighthouse.Router(cloud)
		monitor.Router(cloud)
	}

}
