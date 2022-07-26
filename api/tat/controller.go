package tat

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/member"
)

func listTAT(c *gin.Context) {
	userId := c.GetUint("UserId")
	if res, err := member.ListTAT(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}
}

func infoTAT(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if res, err := member.InfoTAT(id); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}
}

func createTAT(c *gin.Context) {
	var rq member.CreateTATParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := member.CreateTAT(&rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

func updateTAT(c *gin.Context) {
	var rq member.UpdateTATParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}
	if err := member.UpdateTAT(&rq); err == nil {
		c.Set("Payload", "更新成功")
	} else {
		c.Set("Error", err)
	}
}

func deleteTAT(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := member.DeleteTAT(id); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}
}
