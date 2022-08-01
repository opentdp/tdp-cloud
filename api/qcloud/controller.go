package qcloud

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/secret"
	"tdp-cloud/core/qcloud"
)

func doRequest(c *gin.Context) {

	keyId := c.GetUint("KeyId")
	userId := c.GetUint("UserId")

	secret, err := secret.FetchOne(keyId, userId)

	if err != nil || secret.Id == 0 {
		c.Set("Error", "密钥不存在")
		return
	}

	payload, _ := c.GetRawData()

	var params = &qcloud.Params{
		Service:   c.Param("service"),
		Version:   c.Param("version"),
		Action:    c.Param("action"),
		Payload:   payload,
		Region:    c.Param("region"),
		SecretId:  secret.SecretId,
		SecretKey: secret.SecretKey,
	}

	res, err := qcloud.NewRequest(params)

	if err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}
