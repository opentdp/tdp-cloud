package tat

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/tat"
)

func list(c *gin.Context) {
	userId := c.GetUint("UserId")

	if res, err := tat.List(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}
}

func create(c *gin.Context) {
	var rq tat.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := tat.Create(&rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

func update(c *gin.Context) {
	var rq tat.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}
	if err := tat.Update(&rq); err == nil {
		c.Set("Payload", "更新成功")
	} else {
		c.Set("Error", err)
	}
}

func delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := tat.Delete(id); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}
}
