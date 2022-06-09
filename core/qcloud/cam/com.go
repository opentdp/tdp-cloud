package cam

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"

	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
)

// 创建客户端

func NewClient(ud *midware.Userdata) (*cam.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	return cam.NewClient(credential, "", cpf)

}
