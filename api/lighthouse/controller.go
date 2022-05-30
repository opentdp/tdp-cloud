package lighthouse

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 查询地域列表

func describeRegions(c *gin.Context) {

	credential := common.NewCredential(
		c.Request.Header.Get("secretId"),
		c.Request.Header.Get("secretKey"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	client, _ := lighthouse.NewClient(credential, "", cpf)

	request := lighthouse.NewDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Payload": response.Response,
		})
	}

}

// 查看实例列表

func describeInstances(c *gin.Context) {

	credential := common.NewCredential(
		c.Request.Header.Get("secretId"),
		c.Request.Header.Get("secretKey"),
	)

	region := c.Query("region")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	client, _ := lighthouse.NewClient(credential, region, cpf)

	request := lighthouse.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Payload": response.Response,
		})
	}

}

//查看实例流量包详情

func describeInstancesTrafficPackages(c *gin.Context) {

	credential := common.NewCredential(
		c.Request.Header.Get("secretId"),
		c.Request.Header.Get("secretKey"),
	)

	region := c.Query("region")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	client, _ := lighthouse.NewClient(credential, region, cpf)

	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()
	response, err := client.DescribeInstancesTrafficPackages(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Payload": response.Response,
		})
	}

}
