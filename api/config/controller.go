package config

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/cmd/args"
	"tdp-cloud/module/model/config"
)

// 配置列表

func list(c *gin.Context) {

	var rq *config.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if lst, err := config.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Items": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取配置

func detail(c *gin.Context) {

	var rq *config.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if res, err := config.Fetch(rq); err == nil {
		c.Set("Payload", gin.H{"Item": res})
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
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "添加成功")
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

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
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

	var rq *config.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if err := config.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}

// 获取前端配置

func uiOption(c *gin.Context) {

	option := config.ValuesOf("front")

	option["Registrable"] = config.ValueOf("Registrable")
	option["Version"] = args.Version

	c.Set("Payload", option)

}
