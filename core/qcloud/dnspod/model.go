package dnspod

import (
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/utils"
)

// 获取域名列表

type DescribeDomainListRequestParams = dnspod.DescribeDomainListRequestParams

func DescribeDomainList(ud *midware.Userdata, rq *DescribeDomainListRequestParams) (*dnspod.DescribeDomainListResponse, error) {

	client, _ := NewClient(ud)

	request := dnspod.NewDescribeDomainListRequest()
	request.FromJsonString(utils.ToJsonString(rq))

	return client.DescribeDomainList(request)

}

// 获取解析记录列表

type DescribeRecordListRequestParams = dnspod.DescribeRecordListRequestParams

func DescribeRecordList(ud *midware.Userdata, rq *DescribeRecordListRequestParams) (*dnspod.DescribeRecordListResponse, error) {

	client, _ := NewClient(ud)

	request := dnspod.NewDescribeRecordListRequest()
	request.FromJsonString(utils.ToJsonString(rq))

	return client.DescribeRecordList(request)

}

// 修改解析记录

type ModifyRecordRequestParams = dnspod.ModifyRecordRequestParams

func ModifyRecord(ud *midware.Userdata, rq *ModifyRecordRequestParams) (*dnspod.ModifyRecordResponse, error) {

	client, _ := NewClient(ud)

	request := dnspod.NewModifyRecordRequest()
	request.FromJsonString(utils.ToJsonString(rq))

	return client.ModifyRecord(request)

}
