package vendor

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/module/dborm"
	"tdp-cloud/module/dborm/vendor"
)

// 厂商列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if lst, err := vendor.FetchAll(userId); err == nil {
		vendor.SecretMask(lst)
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取厂商

func detail(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if res, err := vendor.Fetch(id, userId); err == nil {
		lst := []*dborm.Vendor{res}
		vendor.SecretMask(lst)
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加厂商

func create(c *gin.Context) {

	var rq *vendor.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := vendor.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 修改厂商

func update(c *gin.Context) {

	var rq *vendor.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := vendor.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除厂商

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if err := vendor.Delete(id, userId); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
