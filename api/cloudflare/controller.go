package cloudflare

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/helper/cloudflare"
	"tdp-cloud/module/dborm/vendor"
)

func apiProxy(c *gin.Context) {

	userId := c.GetUint("UserId")
	vendorId := cast.ToUint(c.Param("id"))

	vendor, err := vendor.Fetch(vendorId, userId)

	if err != nil || vendor.Id == 0 {
		c.Set("Error", "厂商不存在")
		return
	}

	// 构造参数

	params := &cloudflare.Params{
		ApiToken: vendor.SecretKey,
	}

	if err := c.ShouldBindJSON(params); err != nil {
		c.Set("Error", err)
		return
	}

	// 发起请求

	var res *cloudflare.Response

	if c.Request.Method == "GET" {
		res, err = cloudflare.Get(params)
	} else {
		res, err = cloudflare.Post(params)
	}

	if err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
