package cam

import (
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"

	"tdp-cloud/core/midware"
)

// 查询账户摘要

func GetAccountSummary(ud *midware.Userdata) (*cam.GetAccountSummaryResponse, error) {

	client, _ := NewClient(ud)

	request := cam.NewGetAccountSummaryRequest()

	return client.GetAccountSummary(request)

}
