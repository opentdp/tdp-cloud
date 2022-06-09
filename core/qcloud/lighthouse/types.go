package lighthouse

import (
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

type DescribeInstancesRequest struct {
	InstanceIds []*string
	Filters     []*lighthouse.Filter
	Offset      *int64
	Limit       *int64
}

type DescribeInstancesTrafficPackagesRequest struct {
	InstanceIds []*string
	Offset      *int64
	Limit       *int64
}
