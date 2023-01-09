package machine

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/strings"
	"tdp-cloud/internal/dborm/machine"
)

// 域名列表

func list(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := machine.FetchAll(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加域名

func create(c *gin.Context) {

	var rq *machine.CreateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if _, err := machine.Create(rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改域名

func update(c *gin.Context) {

	var rq *machine.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := machine.Update(rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除域名

func delete(c *gin.Context) {

	userId := c.GetUint("UserId")
	id := strings.Uint(c.Param("id"))

	if err := machine.Delete(id, userId); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
