package user

import (
	"errors"

	"github.com/google/uuid"

	"tdp-cloud/module/dborm"
	"tdp-cloud/module/dborm/session"
)

// 添加用户

type CreateParam struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.User{
		AppId:    uuid.NewString(),
		Username: post.Username,
		Password: HashPassword(post.Password),
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新用户

type UpdateParam struct {
	Id          uint
	Description string `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.User{Id: post.Id}).
		Updates(dborm.User{
			Description: post.Description,
		})

	return result.Error

}

// 修改密码

type UpdatePasswordParam struct {
	Id          uint
	OldPassword string `binding:"required"`
	NewPassword string `binding:"required"`
}

func UpdatePassword(post *UpdatePasswordParam) error {

	var item *dborm.User

	// 验证账号

	dborm.Db.Where(&dborm.User{Id: post.Id}).First(&item)

	if item.Id == 0 {
		return errors.New("账号错误")
	}
	if !CheckPassword(item.Password, post.OldPassword) {
		return errors.New("密码错误")
	}

	// 更新密码

	item.Password = HashPassword(post.NewPassword)

	result := dborm.Db.Select("Password").Save(&item)

	return result.Error

}

// 获取用户

type FetchParam struct {
	Id       uint
	AppId    string
	Username string
}

func Fetch(post *FetchParam) (*dborm.User, error) {

	var item *dborm.User

	result := dborm.Db.
		Where(&dborm.User{
			Id:       post.Id,
			AppId:    post.AppId,
			Username: post.Username,
		}).
		First(&item)

	// 删除敏感字段
	item.Password = ""

	return item, result.Error

}

// 登录账号

type LoginParam struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type LoginResult struct {
	AppId       string
	Username    string
	Description string
	Token       string
}

func Login(post *LoginParam) (*LoginResult, error) {

	var item *dborm.User

	// 验证账号

	dborm.Db.Preload("Vendors").Where(&dborm.User{Username: post.Username}).First(&item)

	if item.Id == 0 {
		return nil, errors.New("账号错误")
	}
	if !CheckPassword(item.Password, post.Password) {
		return nil, errors.New("密码错误")
	}

	// 创建令牌

	token, _ := session.Create(item.Id)

	// 返回结果

	res := &LoginResult{
		AppId:       item.AppId,
		Username:    item.Username,
		Description: item.Description,
		Token:       token,
	}

	return res, nil

}
