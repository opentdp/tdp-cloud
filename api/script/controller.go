package script

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/module/dborm/script"
)

// 脚本列表

func list(c *gin.Context) {

	rq := &script.FetchAllParam{
		UserId: c.GetUint("UserId"),
	}

	if lst, err := script.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取脚本

func detail(c *gin.Context) {

	rq := &script.FetchParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

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
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
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

	rq.Id = cast.ToUint(c.Param("id"))
	rq.UserId = c.GetUint("UserId")

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if err := script.Update(rq); err == nil {
		c.Set("Message", "更新成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除脚本

func delete(c *gin.Context) {

	rq := &script.DeleteParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if err := script.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
