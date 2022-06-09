package dnspod

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"

	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// 创建客户端

func NewClient(ud *midware.Userdata) (*dnspod.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	return dnspod.NewClient(credential, "", cpf)

}
