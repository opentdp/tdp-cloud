package qcloud

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	"tdp-cloud/core/midware"
)

func NewCredentialProfile(ud *midware.Userdata) (*common.Credential, *profile.ClientProfile) {

	credential := common.NewCredential(ud.SecretId, ud.SecretKey)

	profile := profile.NewClientProfile()

	return credential, profile

}
