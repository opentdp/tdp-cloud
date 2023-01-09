package tat_history

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/strings"
	history "tdp-cloud/internal/dborm/tat_history"
)

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := history.FetchAll(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

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

func update(c *gin.Context) {

	var rq *history.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := history.Update(rq); err == nil {
		c.Set("Payload", "")
	} else {
		c.Set("Error", err)
	}

}

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := strings.Uint(c.Param("id"))

	if err := history.Delete(id, userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
