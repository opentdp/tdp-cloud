package crontab

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/model/cronjob"
	"tdp-cloud/module/crontab"
)

type Controller struct{}

// 计划列表

func (*Controller) list(c *gin.Context) {

	var rq *cronjob.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if lst, err := cronjob.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{
			"Items":   lst,
			"Entries": crontab.GetEntries(lst),
		})
	} else {
		c.Set("Error", err)
	}

}

// 获取计划

func (*Controller) detail(c *gin.Context) {

	var rq *cronjob.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if res, err := cronjob.Fetch(rq); err == nil {
		c.Set("Payload", gin.H{"Item": res})
	} else {
		c.Set("Error", err)
	}

}

// 添加计划

func (*Controller) create(c *gin.Context) {

	var rq *cronjob.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := cronjob.Create(rq); err == nil {
		crontab.NewById(rq.UserId, id)
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改计划

func (*Controller) update(c *gin.Context) {

	var rq *cronjob.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := cronjob.Update(rq); err == nil {
		crontab.RedoById(rq.UserId, rq.Id)
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除计划

func (*Controller) delete(c *gin.Context) {

	var rq *cronjob.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	crontab.UndoById(rq.UserId, rq.Id)

	if err := cronjob.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
