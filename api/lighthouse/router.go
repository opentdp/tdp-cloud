package lighthouse

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/lighthouse")

	{
		rg.GET("/describeRegions", describeRegionsInstances)

		rg.GET("/describeInstances/:region", describeRegionsInstances)
		rg.GET("/describeRegionsInstances", describeRegionsInstances)

		rg.GET("/describeTrafficPackages/:region", describeTrafficPackages)
		rg.GET("/describeRegionsTrafficPackages", describeRegionsTrafficPackages)
	}

}
