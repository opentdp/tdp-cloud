package sshkey

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/module/dborm/sshkey"
)

// 密钥列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if lst, err := sshkey.FetchAll(userId); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 添加密钥

func create(c *gin.Context) {

	var rq *sshkey.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := sshkey.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := cast.ToUint(c.Param("id"))

	if err := sshkey.Delete(id, userId); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
