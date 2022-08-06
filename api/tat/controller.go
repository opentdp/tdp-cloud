package tat

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/tat"
)

func list(c *gin.Context) {
	userId := c.GetUint("UserId")

	if res, err := tat.FetchAll(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}
}

func create(c *gin.Context) {
	var rq tat.CreateParam

	if c.ShouldBind(&rq) != nil {
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

func update(c *gin.Context) {
	var rq tat.UpdateParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}
	if err := tat.Update(&rq); err == nil {
		c.Set("Payload", "更新成功")
	} else {
		c.Set("Error", err)
	}
}

func delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := tat.Delete(id); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}
}

func create_history(c *gin.Context) {
	var rq tat.AddHistoryParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	userId := c.GetUint("UserId")
	keyId := c.GetUint("KeyId")

	rq.KeyId = keyId
	rq.UserId = userId

	if err := tat.AddHistory(&rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}
}

func list_history(c *gin.Context) {
	var rq tat.AddHistoryParam

	userId := c.GetUint("UserId")
	keyId := c.GetUint("KeyId")

	rq.KeyId = keyId
	rq.UserId = userId

	if res, err := tat.FetchHistory(userId, keyId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}
}

func delete_history(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := tat.DeleteHistory(id); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}
}

func update_history(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var rq tat.UpdateHistoryParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.Id = uint(id)

	if err := tat.UpdateHistory(&rq); err == nil {
		c.Set("Payload", "")
	} else {
		c.Set("Error", err)
	}
}
