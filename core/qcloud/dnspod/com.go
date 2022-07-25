package dnspod

import (
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"
)

// 创建客户端

func NewClient(ud *midware.Userdata) (*dnspod.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	return dnspod.NewClient(credential, "", cpf)

}
