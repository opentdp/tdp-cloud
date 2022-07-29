package qcloud

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"
)

func doRequest(c *gin.Context) {

	payload, _ := c.GetRawData()
	userdata := midware.GetUserdata(c)

	var params = &qcloud.Params{
		Service:   c.Param("service"),
		Version:   c.Param("version"),
		Action:    c.Param("action"),
		Payload:   payload,
		Region:    c.Param("region"),
		SecretId:  userdata.SecretId,
		SecretKey: userdata.SecretKey,
	}

	c.ShouldBindQuery(&params)

	p, _ := json.Marshal(params)
	fmt.Println(string(p))

	res, err := qcloud.NewRequest(params)

	if err == nil {
		c.Set("Payload", res.Response)
	} else {
		c.Set("Error", err)
	}

}
