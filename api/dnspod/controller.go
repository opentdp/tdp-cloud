package dnspod

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// 获取域名列表

func describeDomainList(c *gin.Context) {

	credential := common.NewCredential(
		c.Request.Header.Get("secretId"),
		c.Request.Header.Get("secretKey"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := dnspod.NewClient(credential, "", cpf)

	request := dnspod.NewDescribeDomainListRequest()
	response, err := client.DescribeDomainList(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Payload": response.Response,
		})
	}

}
