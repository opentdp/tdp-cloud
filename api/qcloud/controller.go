package qcloud

import (
	"encoding/json"

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
		SecretId:  secret.SecretId,
		SecretKey: secret.SecretKey,
	}

	header := []byte(c.GetHeader("TDP-QCloud"))

	if json.Unmarshal(header, params) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if payload, err := c.GetRawData(); err == nil {
		params.Payload = payload
	}

	if res, err := qcloud.NewRequest(params); err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}
