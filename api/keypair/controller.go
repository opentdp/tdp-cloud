package keypair

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/module/dborm/keypair"
)

// 密钥列表

func list(c *gin.Context) {

	rq := &keypair.FetchAllParam{
		UserId: c.GetUint("UserId"),
	}

	if lst, err := keypair.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Datasets": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取密钥

func detail(c *gin.Context) {

	rq := &keypair.FetchParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if res, err := keypair.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 添加密钥

func create(c *gin.Context) {

	var rq *keypair.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := keypair.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", gin.H{"Id": id})
	} else {
		c.Set("Error", err)
	}

}

// 修改密钥

func update(c *gin.Context) {

	var rq *keypair.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Id = cast.ToUint(c.Param("id"))
	rq.UserId = c.GetUint("UserId")

	if err := keypair.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func delete(c *gin.Context) {

	rq := &keypair.DeleteParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if err := keypair.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
