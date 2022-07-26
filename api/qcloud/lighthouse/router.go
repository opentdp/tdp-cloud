package lighthouse

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/lighthouse")

	{
		rg.POST("/describeRegions", describeRegions)

		rg.POST("/describeInstances/:region", describeInstances)
		rg.POST("/describeInstancesTrafficPackages/:region", describeInstancesTrafficPackages)

		rg.POST("/describeSnapshots/:region", describeSnapshots)

		rg.POST("/describeFirewallRules/:region", describeFirewallRules)
	}

}
