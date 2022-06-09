package cam

import (
	"tdp-cloud/core/midware"

	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
)

// 查询账户摘要

func GetAccountSummary(ud *midware.Userdata) (*cam.GetAccountSummaryResponse, error) {

	client, _ := NewClient(ud)

	request := cam.NewGetAccountSummaryRequest()

	return client.GetAccountSummary(request)

}
