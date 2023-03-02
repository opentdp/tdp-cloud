package cloudflare

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/helper/cloudflare"
	"tdp-cloud/module/model/vendor"
)

func apiProxy(c *gin.Context) {

	rq := &vendor.FetchParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	vendor, err := vendor.Fetch(rq)

	if err != nil || vendor.Id == 0 {
		c.Set("Error", "厂商不存在")
		return
	}

	// 构造参数

	params := &cloudflare.Params{
		Token: vendor.SecretKey,
	}

	if err := c.ShouldBindJSON(params); err != nil {
		c.Set("Error", err)
		return
	}

	// 发起请求

	res, err := cloudflare.Request(params)

	if err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
