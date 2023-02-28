package passport

import (
	"regexp"

	"github.com/gin-gonic/gin"

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

	if m, n := len(rq.Username), len(rq.Password); m < 4 || n < 6 || m+n > 128 {
		c.Set("Error", "用户名或密码长度不符合要求")
		return
	}

	exp := regexp.MustCompile("^[0-9a-zA-Z\u3040-\u309F\u30A0-\u30FF\u4E00-\u9FA5\uF900-\uFA2D]+$")
	if !exp.MatchString(rq.Username) {
		c.Set("Error", "用户名禁止使用特殊字符")
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

	rq.Id = c.GetUint("UserId")
	rq.Level = 0 //防止逃逸

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
