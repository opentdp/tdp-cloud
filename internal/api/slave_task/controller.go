package slave_task

import (
	"strconv"

	"github.com/gin-gonic/gin"

	task "tdp-cloud/internal/dborm/slave_task"
)

// 任务列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := task.FetchAll(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 获取密钥

func detail(c *gin.Context) {

	userId := c.GetUint("UserId")
	id, _ := strconv.Atoi(c.Param("id"))

	if res, err := task.Fetch(uint(id), userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加任务

func create(c *gin.Context) {

	var rq *task.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if _, err := task.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改配置

func update(c *gin.Context) {

	var rq *task.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := task.Update(rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除任务

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := task.Delete(uint(id), userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
