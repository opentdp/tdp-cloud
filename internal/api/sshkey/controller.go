package sshkey

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/strings"
	"tdp-cloud/internal/dborm/sshkey"
)

// 密钥列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := sshkey.FetchAll(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加密钥

func create(c *gin.Context) {

	var rq *sshkey.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if _, err := sshkey.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := strings.Uint(c.Param("id"))

	if err := sshkey.Delete(id, userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
