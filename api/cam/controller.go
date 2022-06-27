package cam

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud/cam"
)

// 获取账号概要信息

func getAccountSummary(c *gin.Context) {

	var ud = midware.GetUserdata(c)

	if res, err := cam.GetAccountSummary(ud); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}
