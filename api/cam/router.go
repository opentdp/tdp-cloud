package cam

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/cam")

	{
		rg.GET("/getAccountSummary", getAccountSummary)
	}

}
