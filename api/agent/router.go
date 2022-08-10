package agent

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/agent")

	//	rg.Use(midware.AuthGuard())

	{
		rg.GET("/", list)
	}

}
