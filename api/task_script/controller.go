package task_script

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	script "tdp-cloud/module/dborm/task_script"
)

// 脚本列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if lst, err := script.FetchAll(userId); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
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

	if _, err := script.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
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

	rq.UserId = c.GetUint("UserId")

	if err := script.Update(rq); err == nil {
		c.Set("Payload", "更新成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除脚本

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if err := script.Delete(id, userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
