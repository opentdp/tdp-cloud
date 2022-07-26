package dnspod

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/dnspod")

	{
		rg.POST("/describeDomainList", describeDomainList)
		rg.POST("/describeRecordList", describeRecordList)

		rg.POST("/describeRecordLineList", describeRecordLineList)
		rg.POST("/describeRecordType", describeRecordType)

		rg.POST("/modifyRecord", modifyRecord)
	}

}
