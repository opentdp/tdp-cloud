package task_history

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	history "tdp-cloud/internal/dborm/task_history"
)

// 任务列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := history.FetchAll(userId); err == nil {
		c.Set("Payload", history.ParseItems(res))
	} else {
		c.Set("Error", err)
	}

}

// 获取任务

func detail(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if res, err := history.Fetch(id, userId); err == nil {
		c.Set("Payload", history.ParseItem(res))
	} else {
		c.Set("Error", err)
	}

}

// 添加任务

func create(c *gin.Context) {

	var rq *history.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if _, err := history.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改配置

func update(c *gin.Context) {

	var rq *history.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := history.Update(rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除任务

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if err := history.Delete(id, userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
