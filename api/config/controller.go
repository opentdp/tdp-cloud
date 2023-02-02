package config

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/dborm/config"
)

// 配置列表

func list(c *gin.Context) {

	if lst, err := config.FetchAll(); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取配置

func detail(c *gin.Context) {

	name := c.Param("name")

	if res, err := config.Fetch(name); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加配置

func create(c *gin.Context) {

	var rq *config.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := config.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 修改配置

func update(c *gin.Context) {

	var rq *config.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := config.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除配置

func delete(c *gin.Context) {

	name := c.Param("name")

	if err := config.Delete(name); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
