package user

import (
	"tdp-cloud/core/dborm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var post UserInput
	var user dborm.User

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	username := post.Username
	password := post.Password

	dborm.Db.First(&user, "username = ?", username)

	if user.ID > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == nil {
			c.Set("Payload", user)
			return
		}
	}

	c.Set("Error", "账号或密码错误")
}

func Register(c *gin.Context) {

	var post UserInput
	var user dborm.User

	if err := c.BindJSON(&post); err != nil {
		c.Set("Error", "表单错误")
		return
	}

	username := post.Username
	password := post.Password

	dborm.Db.First(&user, "username = ?", username)

	if user.ID > 0 {
		c.Set("Error", "账号已被使用")
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	dborm.Db.Create(&dborm.User{Username: username, Password: string(hash)})

	c.Set("Payload", "账号注册成功")

}
