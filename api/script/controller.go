package script

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/model/script"
)

// 脚本列表

func list(c *gin.Context) {

	var rq *script.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if lst, err := script.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取脚本

func detail(c *gin.Context) {

	var rq *script.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if res, err := script.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加脚本

func create(c *gin.Context) {

	var rq *script.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := script.Create(rq); err == nil {
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改脚本

func update(c *gin.Context) {

	var rq *script.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := script.Update(rq); err == nil {
		c.Set("Message", "更新成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除脚本

func delete(c *gin.Context) {

	var rq *script.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := script.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
