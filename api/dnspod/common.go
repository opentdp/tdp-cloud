package dnspod

import (
	"github.com/gin-gonic/gin"

	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"

	"tdp-cloud/api/qcloud"
)

// 创建客户端

func NewClient(c *gin.Context) *dnspod.Client {

	credential, cpf := qcloud.NewCredentialProfile(c)

	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"

	client, _ := dnspod.NewClient(credential, "", cpf)

	return client

}
