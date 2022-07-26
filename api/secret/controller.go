package secret

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/secret"
)

// 添加密钥

func createSecret(c *gin.Context) {

	var rq secret.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := secret.Create(&rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func deleteSecret(c *gin.Context) {

	userId := c.GetUint("UserId")

	id, _ := strconv.Atoi(c.Param("id"))

	if err := secret.Delete(userId, uint(id)); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}

// 密钥列表

func fetchSecrets(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := secret.Find(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
