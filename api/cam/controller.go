package cam

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/cam"
)

// 获取账号概要信息

func getAccountSummary(c *gin.Context) {

	ud := midware.GetUserdata(c)

	response, err := cam.GetAccountSummary(ud)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
