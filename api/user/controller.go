package user

import (
	"tdp-cloud/core/dborm/user"

	"github.com/gin-gonic/gin"
)

// 注册账号

func register(c *gin.Context) {

	var post user.UserInput

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	ok, err := user.Register(post.Username, post.Password)

	c.Set("Payload", ok)
	c.Set("Error", err)

}

// 登录账号

func login(c *gin.Context) {

	var post user.UserInput

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	token, keyid, err := user.Login(post.Username, post.Password)

	c.Set("Payload", gin.H{"token": token, "keyid": keyid})
	c.Set("Error", err)
}

// 添加密钥

func createSecret(c *gin.Context) {

	var post user.SecretInput

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	userId, _ := c.Get("UserId")
	post.UserID = userId.(uint)

	result, err := user.CreateSecret(&post)

	c.Set("Payload", result)
	c.Set("Error", err)

}

// 删除密钥

func deleteSecret(c *gin.Context) {

	id := c.Param("id")

	result, err := user.DeleteSecret(id)

	c.Set("Payload", result)
	c.Set("Error", err)

}

// 密钥列表

func fetchSecrets(c *gin.Context) {

	userId_, _ := c.Get("UserId")

	list, _ := user.FetchSecrets(userId_.(uint))

	c.Set("Payload", list)

}
