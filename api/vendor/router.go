package vendor

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/vendor", list)
		rg.POST("/vendor", create)
		rg.GET("/vendor/:id", detail)
		rg.PATCH("/vendor/:id", update)
		rg.DELETE("/vendor/:id", delete)
	}

}
