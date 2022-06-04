package dnspod

import (
	"github.com/gin-gonic/gin"

	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// 获取域名列表

func describeDomainList(c *gin.Context) {

	client := NewClient(c)

	request := dnspod.NewDescribeDomainListRequest()
	response, err := client.DescribeDomainList(request)

	c.Set("Payload", response.Response)
	c.Set("Error", err)

}
