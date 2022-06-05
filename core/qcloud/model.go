package qcloud

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func NewCredentialProfile(config [3]string) (*common.Credential, *profile.ClientProfile) {

	credential := common.NewCredential(config[0], config[1])

	profile := profile.NewClientProfile()

	return credential, profile

}
