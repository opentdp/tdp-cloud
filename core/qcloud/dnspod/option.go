package dnspod

import (
	"tdp-cloud/core/midware"

	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// 获取等级允许的线路

type DescribeRecordLineListRequestParams = dnspod.DescribeRecordLineListRequestParams

func DescribeRecordLineList(ud *midware.Userdata, rq *DescribeRecordLineListRequestParams) (*dnspod.DescribeRecordLineListResponse, error) {

	client, _ := NewClient(ud)

	request := dnspod.NewDescribeRecordLineListRequest()

	if rq.Domain != nil {
		request.Domain = rq.Domain
	}

	if rq.DomainGrade != nil {
		request.DomainGrade = rq.DomainGrade
	}

	return client.DescribeRecordLineList(request)

}

// 获取等级允许的记录类型

type DescribeRecordTypeRequestParams = dnspod.DescribeRecordTypeRequestParams

func DescribeRecordType(ud *midware.Userdata, rq *DescribeRecordTypeRequestParams) (*dnspod.DescribeRecordTypeResponse, error) {

	client, _ := NewClient(ud)

	request := dnspod.NewDescribeRecordTypeRequest()

	if rq.DomainGrade != nil {
		request.DomainGrade = rq.DomainGrade
	}

	return client.DescribeRecordType(request)

}
