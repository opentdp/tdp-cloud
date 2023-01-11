package vendor

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/internal/dborm/vendor"
)

// 厂商列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := vendor.FetchAll(userId); err == nil {
		re, _ := regexp.Compile(`^(\w{8}).+(\w{8})$`)
		for k, v := range res {
			res[k].SecretId = re.ReplaceAllString(v.SecretId, "$1*******$2")
			res[k].SecretKey = re.ReplaceAllString(v.SecretKey, "$1******$2")
		}
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加厂商

func create(c *gin.Context) {

	var rq *vendor.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if _, err := vendor.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改厂商

func update(c *gin.Context) {

	var rq *vendor.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := vendor.Update(rq); err == nil {
		c.Set("Payload", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除厂商

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if err := vendor.Delete(id, userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
