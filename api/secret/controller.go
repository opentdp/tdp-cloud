package secret

import (
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/secret"
)

// 密钥列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := secret.FetchAll(userId); err == nil {
		re, _ := regexp.Compile(`^(\w{8}).+(\w{8})$`)
		for k, v := range res {
			res[k].SecretId = re.ReplaceAllString(v.SecretId, "$1*******$2")
			res[k].SecretKey = re.ReplaceAllString(v.SecretKey, "$1******$2")
		}
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加密钥

func create(c *gin.Context) {

	var rq *secret.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := secret.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改配置

func update(c *gin.Context) {

	var rq *secret.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := secret.Update(rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")

	id, _ := strconv.Atoi(c.Param("id"))

	if err := secret.Delete(uint(id), userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
