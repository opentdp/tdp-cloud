package qcloud

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/qcloud"
	"tdp-cloud/helper/strings"
	"tdp-cloud/internal/dborm/vendor"
)

func apiProxy(c *gin.Context) {

	userId := c.GetUint("UserId")
	vendorId := strings.Uint(c.GetHeader("TDP-Vendor"))

	vendor, err := vendor.Fetch(vendorId, userId)

	if err != nil || vendor.Id == 0 {
		c.Set("Error", "厂商不存在")
		return
	}

	params := &qcloud.Params{
		SecretId:   vendor.SecretId,
		SecretKey:  vendor.SecretKey,
		RootDomain: "tencentcloudapi.com",
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
