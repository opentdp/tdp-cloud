package passport

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/cmd/args"
	"tdp-cloud/module/model/passport"
	"tdp-cloud/module/model/user"
)

// 注册用户

func register(c *gin.Context) {

	var rq *user.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	// 是否禁止注册
	if !args.Server.Register {
		c.Set("Error", "抱歉，已关闭注册功能")
		return
	}

	// 校验用户信息
	if err := user.CheckUserinfo(rq.Username, rq.Password, rq.Email); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Level = 0 //防止逃逸

	if id, err := user.Create(rq); err == nil {
		c.Set("Payload", gin.H{"Id": id})
		c.Set("Message", "注册成功")
	} else {
		c.Set("Error", err)
	}

}

// 登录账号

func login(c *gin.Context) {

	var rq *passport.LoginParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.IpAddress = c.ClientIP()
	rq.UserAgent = c.Request.UserAgent()

	if res, err := passport.Login(rq); err == nil {
		c.Set("Message", "登录成功")
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 获取资料

func profile(c *gin.Context) {

	rq := &user.FetchParam{
		Id: c.GetUint("UserId"),
	}

	if res, err := user.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// 修改资料

func profile_update(c *gin.Context) {

	var rq *passport.UpdateInfoParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Id = c.GetUint("UserId")
	rq.AppKey = c.GetString("AppKey")

	if err := passport.UpdateInfo(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 统计信息

func summary(c *gin.Context) {

	userId := c.GetUint("UserId")
	res := passport.Summary(userId)

	c.Set("Payload", res)

}
