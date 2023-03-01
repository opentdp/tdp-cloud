package passport

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"tdp-cloud/module/dborm/passport"
	"tdp-cloud/module/dborm/user"
)

// 注册用户

func register(c *gin.Context) {

	var rq *user.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	// 是否禁止注册
	if !viper.GetBool("server.register") {
		c.Set("Error", "抱歉，已关闭注册功能")
		return
	}

	// 校验用户信息
	if err := user.CheckUser(rq.Username, rq.Password, rq.Email); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Level = 0 //防止逃逸

	if id, err := user.Create(rq); err == nil {
		c.Set("Message", "注册成功")
		c.Set("Payload", gin.H{"Id": id})
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

// 修改资料

func updateInfo(c *gin.Context) {

	var rq *user.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Level = 0     //防止逃逸
	rq.Password = "" //禁止修改
	rq.Id = c.GetUint("UserId")

	// 校验用户信息
	if err := user.CheckUser("", "", rq.Email); err != nil {
		c.Set("Error", err)
		return
	}

	if err := user.Update(rq); err == nil {
		c.Set("Message", "修改成功")
	} else {
		c.Set("Error", err)
	}

}

// 修改密码

func updatePassword(c *gin.Context) {

	var rq *passport.UpdatePasswordParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	rq.Id = c.GetUint("UserId")

	// 校验用户信息
	me, _ := user.Fetch(&user.FetchParam{Id: rq.Id})
	if err := user.CheckUser(me.Username, rq.NewPassword, ""); err != nil {
		c.Set("Error", err)
		return
	}

	if err := passport.UpdatePassword(rq); err == nil {
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
