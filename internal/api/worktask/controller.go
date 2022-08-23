package worktask

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/dborm/worktask"
)

// 任务列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := worktask.FetchAll(userId); err == nil {
		c.Set("Payload", worktask.ParseItems(res))
	} else {
		c.Set("Error", err)
	}

}

// 获取任务

func detail(c *gin.Context) {

	userId := c.GetUint("UserId")
	id, _ := strconv.Atoi(c.Param("id"))

	if res, err := worktask.Fetch(uint(id), userId); err == nil {
		c.Set("Payload", worktask.ParseItem(res))
	} else {
		c.Set("Error", err)
	}

}

// 添加任务

func create(c *gin.Context) {

	var rq *worktask.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if _, err := worktask.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改配置

func update(c *gin.Context) {

	var rq *worktask.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := worktask.Update(rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除任务

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := worktask.Delete(uint(id), userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
