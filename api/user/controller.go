package user

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/dborm/user"
)

// 创建账号

func create(c *gin.Context) {

	var rq *user.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if _, err := user.Create(rq); err == nil {
		c.Set("Payload", "注册成功")
	} else {
		c.Set("Error", err)
	}

}

// 获取用户

func detail(c *gin.Context) {

	rq := &user.FetchParam{
		Id: c.GetUint("UserId"),
	}

	if res, err := user.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 登录账号

func login(c *gin.Context) {

	var rq *user.LoginParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := user.Login(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 修改资料

func updateInfo(c *gin.Context) {

	var rq *user.UpdateInfoParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Id = c.GetUint("UserId")

	if err := user.UpdateInfo(rq); err == nil {
		c.Set("Payload", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改密码

func updatePassword(c *gin.Context) {

	var rq *user.UpdatePasswordParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Id = c.GetUint("UserId")

	if err := user.UpdatePassword(rq); err == nil {
		c.Set("Payload", "修改成功")
	} else {
		c.Set("Error", err)
	}

}
