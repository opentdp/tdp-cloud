package workhub

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/workhub", list)
		rg.POST("/workhub/exec", exec)
	}

}

func Socket(wsi *gin.RouterGroup) {

	rg := wsi.Group("/")

	{
		rg.GET("/workhub", register)
		rg.GET("/workhub/:mid", register)
	}

}
