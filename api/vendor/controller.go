package vendor

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/model/vendor"
)

// 厂商列表

func list(c *gin.Context) {

	var rq *vendor.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if lst, err := vendor.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取厂商

func detail(c *gin.Context) {

	var rq *vendor.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if res, err := vendor.Fetch(rq); err == nil {
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
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "添加成功")
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

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
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

	var rq *vendor.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := vendor.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
