package alibaba

import (
	"github.com/gin-gonic/gin"
	"github.com/open-tdp/go-helper/alibaba"
	"github.com/spf13/cast"

	"tdp-cloud/model/vendor"
)

func apiProxy(c *gin.Context) {

	rq := &vendor.FetchParam{
		Id:       cast.ToUint(c.Param("id")),
		UserId:   c.GetUint("UserId"),
		StoreKey: c.GetString("AppKey"),
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

	param := &alibaba.ReqeustParam{
		SecretId:  vd.SecretId,
		SecretKey: vd.SecretKey,
	}

	if err = c.ShouldBind(param); err != nil {
		c.Set("Error", err)
		return
	}

	// 发起请求

	if res, err := alibaba.Request(param); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
