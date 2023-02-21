package certjob

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/module/certbot"
	"tdp-cloud/module/dborm/certjob"
)

// 计划列表

func list(c *gin.Context) {

	rq := &certjob.FetchAllParam{
		UserId: c.GetUint("UserId"),
	}

	if lst, err := certjob.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取计划

func detail(c *gin.Context) {

	rq := &certjob.FetchParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if res, err := certjob.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加计划

func create(c *gin.Context) {

	var rq *certjob.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := certjob.Create(rq); err == nil {
		certbot.NewById(id)
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 修改计划

func update(c *gin.Context) {

	var rq *certjob.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := certjob.Update(rq); err == nil {
		certbot.RedoById(rq.Id)
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除计划

func delete(c *gin.Context) {

	rq := &certjob.DeleteParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if err := certjob.Delete(rq); err == nil {
		certbot.UndoById(rq.Id)
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
