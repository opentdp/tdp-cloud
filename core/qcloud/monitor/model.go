package monitor

import (
	"tdp-cloud/core/midware"

	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
)

// 获取监控数据

type GetMonitorDataRequestParams = monitor.GetMonitorDataRequestParams

func GetMonitorData(ud *midware.Userdata, rq *GetMonitorDataRequestParams) (*monitor.GetMonitorDataResponse, error) {

	client, _ := NewClient(ud)

	request := monitor.NewGetMonitorDataRequest()

	if rq.Namespace != nil {
		request.Namespace = rq.Namespace
	}

	if rq.MetricName != nil {
		request.MetricName = rq.MetricName
	}

	if rq.Instances != nil {
		request.Instances = rq.Instances
	}

	if rq.Period != nil {
		request.Period = rq.Period
	}

	if rq.StartTime != nil {
		request.StartTime = rq.StartTime
	}

	if rq.EndTime != nil {
		request.EndTime = rq.EndTime
	}

	// request.Namespace = common.StringPtr("QCE/LIGHTHOUSE")
	// request.MetricName = common.StringPtr("LighthouseOuttraffic")

	// request.Instances = []*monitor.Instance{
	// 	&monitor.Instance{
	// 		Dimensions: []*monitor.Dimension{
	// 			&monitor.Dimension{
	// 				Name:  common.StringPtr("InstanceId"),
	// 				Value: common.StringPtr("lhins-noi43mi5"),
	// 			},
	// 		},
	// 	},
	// }

	return client.GetMonitorData(request)

}
