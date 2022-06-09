package qcloud

import (
	"tdp-cloud/core/midware"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func NewCredentialProfile(ud midware.Userdata) (*common.Credential, *profile.ClientProfile) {

	credential := common.NewCredential(ud.SecretId, ud.SecretKey)

	profile := profile.NewClientProfile()

	return credential, profile

}
