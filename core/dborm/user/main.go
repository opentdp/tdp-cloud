package user

import (
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/utils"

	"golang.org/x/crypto/bcrypt"
)

// 注册账号

func Register(username, password string) error {

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	result := dborm.Db.Create(&dborm.User{Username: username, Password: string(hash)})

	return result.Error

}

// 登录账号

func Login(username, password string) (string, int, string) {

	var user dborm.User
	var secret dborm.Secret

	// 验证账号

	dborm.Db.First(&user, "username = ?", username)

	if user.Id == 0 {
		return "", 0, "账号错误"
	}

	// 验证密码

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", 0, "密码错误"
	}

	// 创建令牌

	token := utils.RandString(32)
	dborm.Db.Create(&dborm.Session{UserId: user.Id, Token: token})

	// 获取密钥

	dborm.Db.First(&secret, "user_id = ?", user.Id)

	return token, secret.Id, ""

}
