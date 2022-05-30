package cam

import (
	"github.com/gin-gonic/gin"

	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"

	"tdp-cloud/core/qcloud"
)

// 获取账号概要信息

func getAccountSummary(c *gin.Context) {

	client := qcloud.NewCamClient(c)

	request := cam.NewGetAccountSummaryRequest()
	response, err := client.GetAccountSummary(request)

	c.Set("Payload", response.Response)
	c.Set("Error", err)

}
