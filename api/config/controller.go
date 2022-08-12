package config

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/config"
)

// 配置列表

func list(c *gin.Context) {

	if res, err := config.FetchAll(); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 获取配置

func detail(c *gin.Context) {

	key := c.Param("key")

	if res, err := config.Fetch(key); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加配置

func create(c *gin.Context) {

	var rq *config.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if err := config.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改配置

func update(c *gin.Context) {

	var rq *config.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if err := config.Update(rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除配置

func delete(c *gin.Context) {

	key := c.Param("key")

	if err := config.Delete(key); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
