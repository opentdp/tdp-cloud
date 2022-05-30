package cam

import (
	"net/http"

	"github.com/gin-gonic/gin"

	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func getAccountSummary(c *gin.Context) {

	credential := common.NewCredential(
		c.Request.Header.Get("secretId"),
		c.Request.Header.Get("secretKey"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cam.tencentcloudapi.com"
	client, _ := cam.NewClient(credential, "", cpf)

	request := cam.NewGetAccountSummaryRequest()
	response, err := client.GetAccountSummary(request)

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
