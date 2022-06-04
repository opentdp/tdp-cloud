package lighthouse

import (
	"sync"

	"github.com/gin-gonic/gin"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/api/qcloud"
)

// 创建客户端

func NewClient(c *gin.Context, region string) *lighthouse.Client {

	credential, cpf := qcloud.NewCredentialProfile(c)

	if region == "" {
		cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "lighthouse." + region + ".tencentcloudapi.com"
	}

	client, _ := lighthouse.NewClient(credential, region, cpf)

	return client

}

// 获取所有地域

func DescribeRegions(c *gin.Context) []*lighthouse.RegionInfo {

	client := NewClient(c, "")

	request := lighthouse.NewDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)

	if err == nil && response.Response.RegionSet != nil {
		return response.Response.RegionSet
	}

	return nil

}

// 获取所有实例

func DescribeInstances(c *gin.Context, regionSet []*lighthouse.RegionInfo) []*lighthouse.Instance {

	var wg sync.WaitGroup
	var instanceSet []*lighthouse.Instance

	for _, region := range regionSet {
		wg.Add(1)

		go func(r string) {
			client := NewClient(c, r)

			request := lighthouse.NewDescribeInstancesRequest()
			response, err := client.DescribeInstances(request)

			if err == nil && response.Response.InstanceSet != nil {
				instanceSet = append(instanceSet, response.Response.InstanceSet...)
			}

			wg.Done()
		}(*region.Region)
	}

	wg.Wait()

	return instanceSet

}
