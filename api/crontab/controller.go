package crontab

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/module/crontab"
	"tdp-cloud/module/dborm/cronjob"
)

// 计划列表

func list(c *gin.Context) {

	rq := &cronjob.FetchAllParam{
		UserId: c.GetUint("UserId"),
	}

	if lst, err := cronjob.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取计划

func detail(c *gin.Context) {

	rq := &cronjob.FetchParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if res, err := cronjob.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加计划

func create(c *gin.Context) {

	var rq *cronjob.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := cronjob.Create(rq); err == nil {
		crontab.NewById(id)
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 修改计划

func update(c *gin.Context) {

	var rq *cronjob.UpdateParam

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

	if err := cronjob.Update(rq); err == nil {
		crontab.RedoById(rq.Id)
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除计划

func delete(c *gin.Context) {

	rq := &cronjob.DeleteParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	crontab.UndoById(rq.Id)

	if err := cronjob.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
