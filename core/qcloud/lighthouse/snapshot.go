package lighthouse

import (
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/utils"
)

// 查看快照列表

type DescribeSnapshotsRequestParams = lighthouse.DescribeSnapshotsRequestParams

func DescribeSnapshots(ud *midware.Userdata, rq *DescribeSnapshotsRequestParams) (*lighthouse.DescribeSnapshotsResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeSnapshotsRequest()
	request.FromJsonString(utils.ToJsonString(rq))

	return client.DescribeSnapshots(request)

}
