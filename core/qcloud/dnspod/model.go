package dnspod

import (
	"tdp-cloud/core/midware"

	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// 获取域名列表

type DescribeDomainListRequestParams = dnspod.DescribeDomainListRequestParams

func DescribeDomainList(ud *midware.Userdata, rq *DescribeDomainListRequestParams) (*dnspod.DescribeDomainListResponse, error) {

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

// 获取解析记录列表

type DescribeRecordListRequestParams = dnspod.DescribeRecordListRequestParams

func DescribeRecordList(ud *midware.Userdata, rq *DescribeRecordListRequestParams) (*dnspod.DescribeRecordListResponse, error) {

	client, _ := NewClient(ud)

	request := dnspod.NewDescribeRecordListRequest()

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

	return client.DescribeRecordList(request)

}

// 修改解析记录

type ModifyRecordRequestParams = dnspod.ModifyRecordRequestParams

func ModifyRecord(ud *midware.Userdata, rq *ModifyRecordRequestParams) (*dnspod.ModifyRecordResponse, error) {

	client, _ := NewClient(ud)

	request := dnspod.NewModifyRecordRequest()

	if rq.Domain != nil {
		request.Domain = rq.Domain
	}

	if rq.RecordId != nil {
		request.RecordId = rq.RecordId
	}

	if rq.RecordType != nil {
		request.RecordType = rq.RecordType
	}

	if rq.RecordLine != nil {
		request.RecordLine = rq.RecordLine
	}

	if rq.Value != nil {
		request.Value = rq.Value
	}

	return client.ModifyRecord(request)

}
