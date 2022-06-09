package dnspod

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"

	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// 创建客户端

func NewClient(ud *midware.Userdata) (*dnspod.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	return dnspod.NewClient(credential, "", cpf)

}

// 获取域名列表

func DescribeDomainList(ud *midware.Userdata, rq *DescribeDomainListRequest) (*dnspod.DescribeDomainListResponse, error) {

	client, _ := NewClient(ud)

	request := dnspod.NewDescribeDomainListRequest()

	if rq.Type != nil {
		request.Type = rq.Type
	}

	if rq.Offset != nil {
		request.Offset = rq.Offset
	}

	if rq.Limit != nil {
		request.Limit = rq.Limit
	}

	if rq.GroupId != nil {
		request.GroupId = rq.GroupId
	}

	if rq.Keyword != nil {
		request.Keyword = rq.Keyword
	}

	return client.DescribeDomainList(request)

}
