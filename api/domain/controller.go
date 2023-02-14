package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/module/dborm/domain"
)

// 域名列表

func list(c *gin.Context) {

	rq := &domain.FetchAllParam{
		UserId: c.GetUint("UserId"),
	}

	if lst, err := domain.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取域名

func detail(c *gin.Context) {

	rq := &domain.FetchParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if res, err := domain.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加域名

func create(c *gin.Context) {

	var rq *domain.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := domain.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 修改域名

func update(c *gin.Context) {

	var rq *domain.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := domain.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除域名

func delete(c *gin.Context) {

	rq := &domain.DeleteParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if err := domain.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
