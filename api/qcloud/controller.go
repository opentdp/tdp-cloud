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

	params := &qcloud.Params{
		Service:   c.GetHeader("X-TC-Service"),
		Version:   c.GetHeader("X-TC-Version"),
		Action:    c.GetHeader("X-TC-Action"),
		Region:    c.GetHeader("X-TC-Region"),
		Endpoint:  c.GetHeader("X-TC-Endpoint"),
		SecretId:  secret.SecretId,
		SecretKey: secret.SecretKey,
	}

	if payload, err := c.GetRawData(); err != nil {
		params.Payload = payload
	}

	if res, err := qcloud.NewRequest(params); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}
