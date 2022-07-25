package cam

import (
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"
)

// 创建客户端

func NewClient(ud *midware.Userdata) (*cam.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	return cam.NewClient(credential, "", cpf)

}
