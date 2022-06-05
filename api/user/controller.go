package user

import (
	"tdp-cloud/core/dborm/user"

	"github.com/gin-gonic/gin"
)

// 登录账号

func Login(c *gin.Context) {

	var post UserInput

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	ok, err := user.Login(post.Username, post.Password)

	c.Set("Payload", ok)
	c.Set("Error", err)
}

// 注册账号

func Register(c *gin.Context) {

	var post UserInput

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	ok, err := user.Register(post.Username, post.Password)

	c.Set("Payload", ok)
	c.Set("Error", err)

}
