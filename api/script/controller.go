package script

import (
	"strings"

	"github.com/gin-gonic/gin"

	"tdp-cloud/model/script"
)

type Controller struct{}

// 脚本列表

func (*Controller) list(c *gin.Context) {

	var rq *script.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if lst, err := script.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Items": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取脚本

func (*Controller) detail(c *gin.Context) {

	var rq *script.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if res, err := script.Fetch(rq); err == nil {
		c.Set("Payload", gin.H{"Item": res})
	} else {
		c.Set("Error", err)
	}

}

// 添加脚本

func (*Controller) create(c *gin.Context) {

	var rq *script.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")
	rq.Content = strings.TrimSpace(rq.Content)

	if id, err := script.Create(rq); err == nil {
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改脚本

func (*Controller) update(c *gin.Context) {

	var rq *script.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")
	rq.Content = strings.TrimSpace(rq.Content)

	if err := script.Update(rq); err == nil {
		c.Set("Message", "更新成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除脚本

func (*Controller) delete(c *gin.Context) {

	var rq *script.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := script.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
