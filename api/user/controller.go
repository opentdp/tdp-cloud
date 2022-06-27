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
		c.Set("Error", "请求参数错误")
		return
	}

	if err := member.Register(&rq); err == nil {
		c.Set("Payload", "注册成功")
	} else {
		c.Set("Error", err)
	}

}

// 登录账号

func login(c *gin.Context) {

	var rq member.LoginParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := member.Login(&rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 修改资料

func updateInfo(c *gin.Context) {

	var rq member.UpdateInfoParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := member.UpdateInfo(&rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改密码

func updatePassword(c *gin.Context) {

	var rq member.UpdatePasswordParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := member.UpdatePassword(&rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 添加密钥

func createSecret(c *gin.Context) {

	var rq member.SecretParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := member.CreateSecret(&rq); err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func deleteSecret(c *gin.Context) {

	userId := c.GetUint("UserId")

	id, _ := strconv.Atoi(c.Param("id"))

	if err := member.DeleteSecret(userId, uint(id)); err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}

// 密钥列表

func fetchSecrets(c *gin.Context) {

	userId := c.GetUint("UserId")

	if res, err := member.FindSecrets(userId); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}
