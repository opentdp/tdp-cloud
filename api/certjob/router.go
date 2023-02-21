package certjob

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/certjob", list)
		rg.POST("/certjob", create)
		rg.GET("/certjob/:id", detail)
		rg.PATCH("/certjob/:id", update)
		rg.DELETE("/certjob/:id", delete)
	}

}
