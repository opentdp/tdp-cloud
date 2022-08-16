package slave_node

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/slave")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/node", list)
		rg.POST("/node/exec", exec)
	}

}
