package qcloud

import (
	"github.com/gin-gonic/gin"

	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

func NewCamClient(c *gin.Context) *cam.Client {

	credential, cpf := NewCredentialProfile(c)

	cpf.HttpProfile.Endpoint = "cam.tencentcloudapi.com"

	client, _ := cam.NewClient(credential, "", cpf)

	return client

}

func NewDnspodClient(c *gin.Context) *dnspod.Client {

	credential, cpf := NewCredentialProfile(c)

	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"

	client, _ := dnspod.NewClient(credential, "", cpf)

	return client

}

func NewLighthouseClient(c *gin.Context) *lighthouse.Client {

	region := c.Param("region")

	credential, cpf := NewCredentialProfile(c)

	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"

	client, _ := lighthouse.NewClient(credential, region, cpf)

	return client

}
