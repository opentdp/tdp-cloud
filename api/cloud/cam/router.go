package cam

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/cam")

	rg.Use(midware.Auth())
	rg.Use(midware.Secret())

	{
		rg.POST("/getAccountSummary", getAccountSummary)
	}

}
