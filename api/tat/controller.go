package tat

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/tat"
)

func listTAT(c *gin.Context) {
	userId := c.GetUint("UserId")
	if res, err := tat.List(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}
}

func infoTAT(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if res, err := tat.Info(id); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}
}

func createTAT(c *gin.Context) {
	var rq tat.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := tat.Create(&rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

func updateTAT(c *gin.Context) {
	var rq tat.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}
	if err := tat.Update(&rq); err == nil {
		c.Set("Payload", "更新成功")
	} else {
		c.Set("Error", err)
	}
}

func deleteTAT(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := tat.Delete(id); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}
}
