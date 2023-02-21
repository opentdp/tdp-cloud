package cronjob

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/cronjob", list)
		rg.POST("/cronjob", create)
		rg.GET("/cronjob/:id", detail)
		rg.PATCH("/cronjob/:id", update)
		rg.DELETE("/cronjob/:id", delete)
	}

}
