package machine

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/module/dborm/machine"
)

// 主机列表

func list(c *gin.Context) {

	rq := &machine.FetchAllParam{
		UserId: c.GetUint("UserId"),
	}

	if lst, err := machine.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取主机

func detail(c *gin.Context) {

	rq := &machine.FetchParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if res, err := machine.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加主机

func create(c *gin.Context) {

	var rq *machine.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := machine.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 修改主机

func update(c *gin.Context) {

	var rq *machine.UpdateParam

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

	if err := machine.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除主机

func delete(c *gin.Context) {

	rq := &machine.DeleteParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if err := machine.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
