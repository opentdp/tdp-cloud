package user

import (
	"tdp-cloud/core/dborm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 登录账号

func Login(c *gin.Context) {

	var post UserInput
	var user dborm.User

	// 验证表单

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	username := post.Username
	password := post.Password

	// 验证账号

	dborm.Db.First(&user, "username = ?", username)

	if user.ID == 0 {
		c.Set("Error", "账号错误")
		return
	}

	// 验证密码

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		c.Set("Error", "密码错误")
		return
	}

	// 创建令牌

	token := RandString(32)
	dborm.Db.Create(&dborm.Session{UserID: user.ID, Token: token})

	c.Set("Payload", token)
}

// 注册账号

func Register(c *gin.Context) {

	var post UserInput
	var user dborm.User

	// 验证表单

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	username := post.Username
	password := post.Password

	// 验证账号

	dborm.Db.First(&user, "username = ?", username)

	if user.ID > 0 {
		c.Set("Error", "账号已被使用")
		return
	}

	// 创建账号

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	dborm.Db.Create(&dborm.User{Username: username, Password: string(hash)})

	c.Set("Payload", "账号注册成功")

}
