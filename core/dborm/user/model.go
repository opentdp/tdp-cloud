package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"tdp-cloud/core/dborm"
	"tdp-cloud/core/dborm/session"
)

// 注册账号

type RegisterParam struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func Register(param *RegisterParam) error {

	hash, _ := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	result := dborm.Db.Create(&dborm.User{Username: param.Username, Password: string(hash)})

	return result.Error

}

// 登录账号

type LoginParam struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type LoginResult struct {
	KeyId       uint
	Token       string
	Username    string
	Description string
}

func Login(param *LoginParam) (*LoginResult, error) {

	var item *dborm.User

	// 验证账号

	dborm.Db.Preload("Secrets").First(&item, "username = ?", param.Username)

	if item.Id == 0 {
		return nil, errors.New("账号错误")
	}

	// 验证密码

	err := bcrypt.CompareHashAndPassword([]byte(item.Password), []byte(param.Password))

	if err != nil {
		return nil, errors.New("密码错误")
	}

	// 创建令牌

	token, _ := session.Create(item.Id)

	// 返回结果

	res := &LoginResult{
		Token:       token,
		Username:    item.Username,
		Description: item.Description,
	}

	if len(item.Secrets) > 0 {
		res.KeyId = item.Secrets[0].Id
	}

	return res, nil

}

// 修改资料

type UpdateInfoParam struct {
	UserId      uint
	Description string `binding:"required"`
}

func UpdateInfo(param *UpdateInfoParam) error {

	var item *dborm.User

	// 验证账号

	dborm.Db.First(&item, "id = ?", param.UserId)

	if item.Id == 0 {
		return errors.New("账号错误")
	}

	// 更新资料

	item.Description = param.Description

	dborm.Db.Select("Description").Save(&item)

	return nil

}

// 修改密码

type UpdatePasswordParam struct {
	UserId      uint
	OldPassword string `binding:"required"`
	NewPassword string `binding:"required"`
}

func UpdatePassword(param *UpdatePasswordParam) error {

	var item *dborm.User

	// 验证账号

	dborm.Db.First(&item, "id = ?", param.UserId)

	if item.Id == 0 {
		return errors.New("账号错误")
	}

	err := bcrypt.CompareHashAndPassword([]byte(item.Password), []byte(param.OldPassword))

	if err != nil {
		return errors.New("密码错误")
	}

	// 更新密码

	hash, _ := bcrypt.GenerateFromPassword([]byte(param.NewPassword), bcrypt.DefaultCost)
	item.Password = string(hash)

	dborm.Db.Select("Password").Save(&item)

	return nil

}
