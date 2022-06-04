package lighthouse

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/lighthouse")

	{
		rg.GET("/getAllRegionsInstances", getAllRegionsInstances)
		rg.GET("/describeInstancesTrafficPackages/:region", describeInstancesTrafficPackages)
	}

}
