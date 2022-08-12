package qcloud

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/secret"
	"tdp-cloud/core/qcloud"
)

func apiProxy(c *gin.Context) {

	keyId := c.GetUint("KeyId")
	userId := c.GetUint("UserId")

	secret, err := secret.Fetch(keyId, userId)

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
		re, _ := regexp.Compile(`^.+, Message=`)
		msg := re.ReplaceAllString(err.Error(), "")
		c.Set("Error", msg)
	}

}

func vncProxy(c *gin.Context) {

	resp, err := http.Get("https://img.qcloud.com/qcloud/app/active_vnc/index.html")

	if err != nil {
		c.Set("Error", "获取资源失败")
		return
	}

	if res, err := ioutil.ReadAll(resp.Body); err == nil {
		c.Set("HTML", string(res))
	} else {
		c.Set("Error", err)
	}

}
