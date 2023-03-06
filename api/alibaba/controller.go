package alibaba

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/helper/alibaba"
	"tdp-cloud/module/model/vendor"
)

func apiProxy(c *gin.Context) {

	rq := &vendor.FetchParam{
		Id:       cast.ToUint(c.Param("id")),
		UserId:   c.GetUint("UserId"),
		StoreKey: c.GetString("appkey"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	vd, err := vendor.Fetch(rq)

	if err != nil || vd.Id == 0 {
		c.Set("Error", "厂商不存在")
		return
	}

	// 构造参数

	params := &alibaba.Params{
		SecretId:  vd.SecretId,
		SecretKey: vd.SecretKey,
	}

	if err = c.ShouldBindJSON(params); err != nil {
		c.Set("Error", err)
		return
	}

	// 发起请求

	if res, err := alibaba.Request(params); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
