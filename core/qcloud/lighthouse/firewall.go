package lighthouse

import (
	"tdp-cloud/core/midware"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 查询防火墙规则

type DescribeFirewallRulesRequest struct {
	InstanceId *string
	Offset     *int64
	Limit      *int64
}

func DescribeFirewallRules(ud *midware.Userdata, rq *DescribeFirewallRulesRequest) (*lighthouse.DescribeFirewallRulesResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeFirewallRulesRequest()

	if rq.InstanceId != nil {
		request.InstanceId = rq.InstanceId
	}

	if rq.Offset != nil {
		request.Offset = rq.Offset
	}

	if rq.Limit != nil {
		request.Limit = rq.Limit
	}

	return client.DescribeFirewallRules(request)

}
