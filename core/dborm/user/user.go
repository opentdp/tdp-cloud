package user

import (
	"errors"
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

type LoginResult struct {
	Keyid int    `json:"keyid"`
	Token string `json:"token"`
}

func Login(username, password string) (LoginResult, error) {

	var res LoginResult

	var user dborm.User
	var secret dborm.Secret

	// 验证账号

	dborm.Db.First(&user, "username = ?", username)

	if user.Id == 0 {
		return res, errors.New("账号错误")
	}

	// 验证密码

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return res, errors.New("密码错误")
	}

	// 创建令牌

	token := utils.RandString(32)
	dborm.Db.Create(&dborm.Session{UserId: user.Id, Token: token})

	// 获取密钥

	dborm.Db.First(&secret, "user_id = ?", user.Id)

	res.Keyid = secret.Id
	res.Token = token

	return res, nil

}

// 修改资料

type ModifyQuery struct {
	UserId      int    `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Description string `json:"description"`
}

func Modify(data ModifyQuery) error {

	var user dborm.User

	// 验证账号

	dborm.Db.First(&user, "user_id = ?", data.UserId)

	if user.Id == 0 {
		return errors.New("账号错误")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.OldPassword))

	if err != nil {
		return errors.New("密码错误")
	}

	// 更新资料

	if data.NewPassword != "" {
		password, _ := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
		user.Password = string(password)
	}

	user.Description = data.Description

	dborm.Db.Save(&user)

	return nil

}
