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

func Login(data *LoginParam) (*LoginResult, error) {

	item, _ := user.Fetch(&user.FetchParam{
		Username: data.Username,
	})

	// 验证账号

	if item.Id == 0 {
		return nil, errors.New("账号错误")
	}
	if !user.CheckPassword(item.Password, data.Password) {
		return nil, errors.New("密码错误")
	}

	// 创建令牌

	token, _ := session.Create(&session.CreateParam{
		UserId:    item.Id,
		UserAgent: "",
	})

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

func UpdatePassword(data *UpdatePasswordParam) error {

	item, _ := user.Fetch(&user.FetchParam{Id: data.Id})

	// 验证账号

	if item.Id == 0 {
		return errors.New("账号错误")
	}
	if !user.CheckPassword(item.Password, data.OldPassword) {
		return errors.New("密码错误")
	}

	// 更新密码

	return user.Update(&user.UpdateParam{
		Id:       data.Id,
		Password: data.NewPassword,
	})

}
