package lighthouse

import (
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/utils"
)

// 查询防火墙规则

type DescribeFirewallRulesRequestParams = lighthouse.DescribeFirewallRulesRequestParams

func DescribeFirewallRules(ud *midware.Userdata, rq *DescribeFirewallRulesRequestParams) (*lighthouse.DescribeFirewallRulesResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeFirewallRulesRequest()
	request.FromJsonString(utils.ToJsonString(rq))

	return client.DescribeFirewallRules(request)

}
