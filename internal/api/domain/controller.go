package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/internal/dborm/domain"
)

// 域名列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := domain.FetchAll(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 获取域名

func detail(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if res, err := domain.Fetch(id, userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加域名

func create(c *gin.Context) {

	var rq *domain.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if _, err := domain.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改域名

func update(c *gin.Context) {

	var rq *domain.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := domain.Update(rq); err == nil {
		c.Set("Payload", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除域名

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if err := domain.Delete(id, userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
