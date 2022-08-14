package user

import (
	"errors"

	"github.com/google/uuid"

	"tdp-cloud/core/dborm"
	"tdp-cloud/core/dborm/session"
)

// 创建账号

type CreateParam struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func Create(param *CreateParam) error {

	result := dborm.Db.Create(&dborm.User{
		Username: param.Username,
		Password: HashPassword(param.Password),
		AppToken: uuid.NewString(),
	})

	return result.Error

}

// 修改资料

type UpdateInfoParam struct {
	Id          uint
	Description string `binding:"required"`
}

func UpdateInfo(param *UpdateInfoParam) error {

	var item *dborm.User

	// 验证账号

	dborm.Db.First(&item, "id = ?", param.Id)

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
	Id          uint
	OldPassword string `binding:"required"`
	NewPassword string `binding:"required"`
}

func UpdatePassword(param *UpdatePasswordParam) error {

	var item *dborm.User

	// 验证账号

	dborm.Db.First(&item, "id = ?", param.Id)

	if item.Id == 0 {
		return errors.New("账号错误")
	}
	if !CheckPassword(item.Password, param.OldPassword) {
		return errors.New("密码错误")
	}

	// 更新密码

	item.Password = HashPassword(param.NewPassword)

	dborm.Db.Select("Password").Save(&item)

	return nil

}

// 获取用户

type FetchParam struct {
	Id       uint
	Username string
	AppToken string
}

func Fetch(param *FetchParam) (*dborm.User, error) {

	var item *dborm.User

	dborm.Db.Where(param).First(&item)

	if item.Id == 0 {
		return nil, errors.New("用户不存在")
	}

	// 删除敏感字段
	item.Password = ""

	return item, nil

}

// 登录账号

type LoginParam struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type LoginResult struct {
	KeyId        uint
	AppToken     string
	SessionToken string
	Username     string
	Description  string
}

func Login(param *LoginParam) (*LoginResult, error) {

	var item *dborm.User

	// 验证账号

	dborm.Db.Preload("Secrets").First(&item, "username = ?", param.Username)

	if item.Id == 0 {
		return nil, errors.New("账号错误")
	}
	if !CheckPassword(item.Password, param.Password) {
		return nil, errors.New("密码错误")
	}

	// 创建令牌

	token, _ := session.Create(item.Id)

	// 返回结果

	res := &LoginResult{
		Username:     item.Username,
		AppToken:     item.AppToken,
		Description:  item.Description,
		SessionToken: token,
	}

	if len(item.Secrets) > 0 {
		res.KeyId = item.Secrets[0].Id
	}

	return res, nil

}
