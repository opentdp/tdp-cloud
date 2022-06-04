package dnspod

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/dnspod")

	{
		rg.GET("/describeDomainList", describeDomainList)
	}

}
