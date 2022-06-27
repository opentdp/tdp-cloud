package user

import (
	"strconv"
	"tdp-cloud/core/dborm/member"

	"github.com/gin-gonic/gin"
)

// 注册账号

func register(c *gin.Context) {

	var rq member.RegisterParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	err := member.Register(&rq)

	if err == nil {
		c.Set("Payload", "注册成功")
	} else {
		c.Set("Error", err)
	}

}

// 登录账号

func login(c *gin.Context) {

	var rq member.LoginParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	res, err := member.Login(&rq)

	if err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 修改资料

func updateInfo(c *gin.Context) {

	var rq member.UpdateInfoParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	err := member.UpdateInfo(&rq)

	if err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改密码

func updatePassword(c *gin.Context) {

	var rq member.UpdatePasswordParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	err := member.UpdatePassword(&rq)

	if err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 添加密钥

func createSecret(c *gin.Context) {

	var rq member.SecretParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.UserId = c.GetUint("UserId")

	err := member.CreateSecret(&rq)

	if err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func deleteSecret(c *gin.Context) {

	userId := c.GetUint("UserId")

	id, _ := strconv.Atoi(c.Param("id"))

	err := member.DeleteSecret(userId, uint(id))

	if err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}

// 密钥列表

func fetchSecrets(c *gin.Context) {

	userId := c.GetUint("UserId")

	list := member.FindSecrets(userId)

	c.Set("Payload", list)

}
