package cam

import (
	"github.com/gin-gonic/gin"

	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"

	"tdp-cloud/api/qcloud"
)

// 创建客户端

func NewClient(c *gin.Context) *cam.Client {

	credential, cpf := qcloud.NewCredentialProfile(c)

	cpf.HttpProfile.Endpoint = "cam.tencentcloudapi.com"

	client, _ := cam.NewClient(credential, "", cpf)

	return client

}
