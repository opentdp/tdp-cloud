package lighthouse

import (
	"tdp-cloud/core/midware"

	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/lighthouse")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	{
		rg.GET("/describeRegions", describeRegions)

		rg.GET("/describeInstances/:region", describeInstances)
		rg.GET("/describeInstancesTrafficPackages/:region", describeInstancesTrafficPackages)

		rg.GET("/describeSnapshots/:region", describeSnapshots)

		rg.GET("/describeFirewallRules/:region", describeFirewallRules)
	}

}
