package passport

import (
	"errors"

	"tdp-cloud/module/dborm/session"
	"tdp-cloud/module/dborm/user"
)

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

	item, _ := user.Fetch(&user.FetchParam{Username: post.Username})

	// 验证账号

	if item.Id == 0 {
		return nil, errors.New("账号错误")
	}
	if !user.CheckPassword(item.Password, post.Password) {
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

// 修改密码

type UpdatePasswordParam struct {
	Id          uint
	OldPassword string `binding:"required"`
	NewPassword string `binding:"required"`
}

func UpdatePassword(post *UpdatePasswordParam) error {

	item, _ := user.Fetch(&user.FetchParam{Id: post.Id})

	// 验证账号

	if item.Id == 0 {
		return errors.New("账号错误")
	}
	if !user.CheckPassword(item.Password, post.OldPassword) {
		return errors.New("密码错误")
	}

	// 更新密码

	return user.Update(&user.UpdateParam{
		Id:       post.Id,
		Password: post.NewPassword,
	})

}
