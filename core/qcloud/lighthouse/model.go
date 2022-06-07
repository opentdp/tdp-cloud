package lighthouse

import (
	"sync"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/qcloud"
)

// 创建客户端

func NewClient(config [3]string) (*lighthouse.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(config)

	if config[2] == "" {
		cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "lighthouse." + config[2] + ".tencentcloudapi.com"
	}

	client, err := lighthouse.NewClient(credential, config[2], cpf)

	return client, err

}

// 获取地域列表

func DescribeRegions(config [3]string) (*lighthouse.DescribeRegionsResponse, error) {

	client, err := NewClient(config)

	if err != nil {
		return nil, err
	}

	request := lighthouse.NewDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)

	return response, nil

}

// 获取实列列表 - 单地域

func DescribeInstances(config [3]string) (*lighthouse.DescribeInstancesResponse, error) {

	client, err := NewClient(config)

	if err != nil {
		return nil, err
	}

	request := lighthouse.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)

	return response, err

}

// 获取实列列表 - 多地域

func DescribeInstancesAll(config [3]string, regionSet []*lighthouse.RegionInfo) ([]*lighthouse.Instance, []error) {

	var ers []error

	var wg sync.WaitGroup
	var instanceSet []*lighthouse.Instance

	for _, region := range regionSet {
		wg.Add(1)

		go func(r string) {
			c := [3]string{config[0], config[1], r}
			response, err := DescribeInstances(c)

			if err != nil {
				ers = append(ers, err)
			} else if response.Response.InstanceSet != nil {
				instanceSet = append(instanceSet, response.Response.InstanceSet...)
			}

			wg.Done()
		}(*region.Region)
	}

	wg.Wait()

	return instanceSet, ers

}

// 查看流量包详情 - 单地域

func DescribeTrafficPackages(config [3]string) (*lighthouse.DescribeInstancesTrafficPackagesResponse, error) {

	client, err := NewClient(config)

	if err != nil {
		return nil, err
	}

	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()
	response, err := client.DescribeInstancesTrafficPackages(request)

	return response, err

}

// 查看流量包详情 - 多地域

func DescribeTrafficPackagesAll(config [3]string, regionSet []*lighthouse.RegionInfo) ([]*lighthouse.InstanceTrafficPackage, []error) {

	var ers []error

	var wg sync.WaitGroup
	var trafficPackageSet []*lighthouse.InstanceTrafficPackage

	for _, region := range regionSet {
		wg.Add(1)

		go func(r string) {
			c := [3]string{config[0], config[1], r}
			response, err := DescribeTrafficPackages(c)

			if err != nil {
				ers = append(ers, err)
			} else if response.Response.InstanceTrafficPackageSet != nil {
				trafficPackageSet = append(trafficPackageSet, response.Response.InstanceTrafficPackageSet...)
			}

			wg.Done()
		}(*region.Region)
	}

	wg.Wait()

	return trafficPackageSet, ers

}
