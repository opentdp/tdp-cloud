package lighthouse

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/lighthouse")

	{
		rg.GET("/describeRegions", describeRegions)
		rg.GET("/describeInstances", describeInstances)
		rg.GET("/describeInstancesTrafficPackages", describeInstancesTrafficPackages)
	}

}
