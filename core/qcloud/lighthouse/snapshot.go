package lighthouse

import (
	"tdp-cloud/core/midware"

	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

// 查看快照列表

type DescribeSnapshotsRequestParams = lighthouse.DescribeSnapshotsRequestParams

func DescribeSnapshots(ud *midware.Userdata, rq *DescribeSnapshotsRequestParams) (*lighthouse.DescribeSnapshotsResponse, error) {

	client, _ := NewClient(ud)

	request := lighthouse.NewDescribeSnapshotsRequest()

	if len(rq.SnapshotIds) > 0 {
		request.SnapshotIds = rq.SnapshotIds
	}

	if len(rq.Filters) > 0 {
		request.Filters = rq.Filters
	}

	if rq.Offset != nil {
		request.Offset = rq.Offset
	}

	if rq.Limit != nil {
		request.Limit = rq.Limit
	}

	return client.DescribeSnapshots(request)

}
