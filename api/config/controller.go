package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/cmd/args"
	"tdp-cloud/module/model/config"
)

// 配置列表

func list(c *gin.Context) {

	rq := &config.FetchAllParam{
		Module: c.Param("module"),
	}

	if lst, err := config.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取配置

func detail(c *gin.Context) {

	rq := &config.FetchParam{
		Id: cast.ToUint(c.Param("id")),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if res, err := config.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

func detail_name(c *gin.Context) {

	rq := &config.FetchParam{
		Name: c.Param("name"),
	}

	if res, err := config.Fetch(rq); err == nil {
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

	rq.Id = cast.ToUint(c.Param("id"))

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

	rq := &config.DeleteParam{
		Id: cast.ToUint(c.Param("id")),
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

func ui_option(c *gin.Context) {

	option := gin.H{
		"register": args.Server.Register,
	}

	c.Set("Payload", option)

}
