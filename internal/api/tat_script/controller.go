package tat_script

import (
	"strconv"

	"github.com/gin-gonic/gin"

	script "tdp-cloud/internal/dborm/tat_script"
)

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := script.FetchAll(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

func create(c *gin.Context) {

	var rq *script.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if _, err := script.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

func update(c *gin.Context) {

	var rq *script.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := script.Update(rq); err == nil {
		c.Set("Payload", "更新成功")
	} else {
		c.Set("Error", err)
	}

}

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")

	id, _ := strconv.Atoi(c.Param("id"))

	if err := script.Delete(uint(id), userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
