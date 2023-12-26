package alibaba

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/alibaba"
	"github.com/opentdp/go-helper/strutil"

	"tdp-cloud/model/vendor"
)

type Controller struct{}

func (*Controller) apiProxy(c *gin.Context) {

	rq := &vendor.FetchParam{
		Id:       strutil.ToUint(c.Param("id")),
		UserId:   c.GetUint("UserId"),
		StoreKey: c.GetString("AppKey"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	vdr, err := vendor.Fetch(rq)

	if err != nil || vdr.Id == 0 {
		c.Set("Error", "厂商不存在")
		return
	}

	// 构造参数

	param := &alibaba.ReqeustParam{
		SecretId:  vdr.SecretId,
		SecretKey: vdr.SecretKey,
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
