package user

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/core/dborm/user"
)

// 注册账号

func register(c *gin.Context) {

	var rq user.RegisterParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if err := user.Register(&rq); err == nil {
		c.Set("Payload", "注册成功")
	} else {
		c.Set("Error", err)
	}

}

// 登录账号

func login(c *gin.Context) {

	var rq user.LoginParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	if res, err := user.Login(&rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 修改资料

func updateInfo(c *gin.Context) {

	var rq user.UpdateInfoParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := user.UpdateInfo(&rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改密码

func updatePassword(c *gin.Context) {

	var rq user.UpdatePasswordParam

	if c.ShouldBind(&rq) != nil {
		c.Set("Error", "请求参数错误")
		return
	}

	rq.UserId = c.GetUint("UserId")

	if err := user.UpdatePassword(&rq); err == nil {
		c.Set("Payload", "操作成功")
	} else {
		c.Set("Error", err)
	}

}
