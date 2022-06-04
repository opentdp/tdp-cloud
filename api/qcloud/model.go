package qcloud

import (
	"github.com/gin-gonic/gin"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func NewCredentialProfile(c *gin.Context) (*common.Credential, *profile.ClientProfile) {

	credential := common.NewCredential(
		c.Request.Header.Get("secretId"),
		c.Request.Header.Get("secretKey"),
	)

	profile := profile.NewClientProfile()

	return credential, profile

}
