package lighthouse

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/lighthouse")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	{
		rg.POST("/describeRegions", describeRegions)

		rg.POST("/describeInstances/:region", describeInstances)
		rg.POST("/describeInstancesTrafficPackages/:region", describeInstancesTrafficPackages)

		rg.POST("/describeSnapshots/:region", describeSnapshots)

		rg.POST("/describeFirewallRules/:region", describeFirewallRules)
	}

}
