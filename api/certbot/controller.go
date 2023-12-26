package certbot

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/model/certjob"
	"tdp-cloud/module/certbot"
)

type Controller struct{}

// 计划列表

func (*Controller) list(c *gin.Context) {

	var rq *certjob.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if lst, err := certjob.FetchAll(rq); err == nil {
		c.Set("Payload", gin.H{"Items": lst})
	} else {
		c.Set("Error", err)
	}

}

// 获取计划

func (*Controller) detail(c *gin.Context) {

	var rq *certjob.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if job, crt, err := certbot.CertById(rq.UserId, rq.Id); err == nil {
		c.Set("Payload", gin.H{"Item": job, "Cert": crt})
	} else {
		c.Set("Error", err)
	}

}

// 添加计划

func (*Controller) create(c *gin.Context) {

	var rq *certjob.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	if id, err := certjob.Create(rq); err == nil {
		certbot.NewById(rq.UserId, id)
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改计划

func (*Controller) update(c *gin.Context) {

	var rq *certjob.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := certjob.Update(rq); err == nil {
		certbot.RedoById(rq.UserId, rq.Id)
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除计划

func (*Controller) delete(c *gin.Context) {

	var rq *certjob.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	certbot.UndoById(rq.UserId, rq.Id)

	if err := certjob.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
