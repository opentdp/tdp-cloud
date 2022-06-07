package lighthouse

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/lighthouse")

	{
		rg.GET("/describeRegionsInstances", describeRegionsInstances)

		rg.GET("/describeInstancesTrafficPackages", DescribeInstancesTrafficPackagesAll)
		rg.GET("/describeInstancesTrafficPackages/:region", describeInstancesTrafficPackages)
	}

}
