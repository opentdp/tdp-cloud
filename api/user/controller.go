package user

import (
	"strconv"
	"tdp-cloud/core/dborm/user"

	"github.com/gin-gonic/gin"
)

// 注册账号

func register(c *gin.Context) {

	var param user.RegisterParam

	if err := c.BindJSON(&param); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	err := user.Register(&param)

	if err == nil {
		c.Set("Payload", "注册成功")
	} else {
		c.Set("Error", err)
	}

}

// 登录账号

func login(c *gin.Context) {

	var param user.LoginParam

	if err := c.BindJSON(&param); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	res, err := user.Login(&param)

	if err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 修改资料

func modify(c *gin.Context) {

	var param user.ModifyParam

	if err := c.BindJSON(&param); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	param.UserId = c.GetInt("UserId")

	err := user.Modify(&param)

	if err == nil {
		c.Set("Payload", "")
	} else {
		c.Set("Error", err)
	}

}

// 添加密钥

func createSecret(c *gin.Context) {

	var param user.SecretParam

	if err := c.BindJSON(&param); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	param.UserId = c.GetInt("UserId")

	err := user.CreateSecret(&param)

	if err == nil {
		c.Set("Payload", "添加成功")
	} else {
		c.Set("Error", err)
	}

}

// 删除密钥

func deleteSecret(c *gin.Context) {

	UserId := c.GetInt("UserId")

	id, _ := strconv.Atoi(c.Param("id"))

	err := user.DeleteSecret(UserId, id)

	if err == nil {
		c.Set("Payload", "删除成功")
	} else {
		c.Set("Error", err)
	}

}

// 密钥列表

func fetchSecrets(c *gin.Context) {

	userId := c.GetInt("UserId")

	list := user.FindSecrets(userId)

	c.Set("Payload", list)

}
