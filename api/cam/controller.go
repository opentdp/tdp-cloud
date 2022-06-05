package cam

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/qcloud/cam"
)

// 获取账号概要信息

func getAccountSummary(c *gin.Context) {

	config_, _ := c.Get("Config")
	config := config_.([3]string)

	response, err := cam.GetAccountSummary(config)

	if response != nil {
		c.Set("Payload", response.Response)
	}

	c.Set("Error", err)

}
