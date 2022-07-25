package monitor

import (
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/utils"
)

// 获取监控数据

type GetMonitorDataRequestParams = monitor.GetMonitorDataRequestParams

func GetMonitorData(ud *midware.Userdata, rq *GetMonitorDataRequestParams) (*monitor.GetMonitorDataResponse, error) {

	client, _ := NewClient(ud)

	request := monitor.NewGetMonitorDataRequest()
	request.FromJsonString(utils.ToJsonString(rq))

	return client.GetMonitorData(request)

}
