package alibaba

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/alibaba")

	// 需授权接口

	rg.Use(midware.AuthGuard())

	{
		rg.POST("/:id", apiProxy)
	}

}
