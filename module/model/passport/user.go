package passport

import (
	"errors"
	"fmt"

	"tdp-cloud/module/midware"
	"tdp-cloud/module/model/user"
)

// 登录账号

type LoginParam struct {
	Username  string `binding:"required"`
	Password  string `binding:"required"`
	IpAddress string
	UserAgent string
}

type LoginResult struct {
	Username string
	AppId    string
	Email    string
	Token    string
}

func Login(data *LoginParam) (*LoginResult, error) {

	item, _ := user.Fetch(&user.FetchParam{
		Username: data.Username,
	})

	// 验证账号

	fmt.Println(item.Password, data.Password)

	if item.Id == 0 {
		return nil, errors.New("账号错误")
	}
	if !user.CheckSecret(item.Password, data.Password) {
		return nil, errors.New("密码错误")
	}

	// 自动迁移密钥
	// TODO: v1.0.0 时删除兼容代码

	// item.Password = data.Password
	if temp, err := secretMigrator(item); err != nil {
		return nil, err
	} else {
		item = temp
	}

	// 创建令牌

	token, err := midware.CreateToken(&midware.UserInfo{
		AppKey:    item.AppKey,
		UserId:    item.Id,
		UserLevel: item.Level,
	})

	if err != nil {
		return nil, err
	}

	// 返回结果

	res := &LoginResult{
		Username: item.Username,
		AppId:    item.AppId,
		Email:    item.Email,
		Token:    token,
	}

	return res, nil

}

// 修改资料

type ProfileUpdateParam struct {
	user.UpdateParam
	OldPassword string `binding:"required"`
}

func ProfileUpdate(data *ProfileUpdateParam) error {

	item, _ := user.Fetch(&user.FetchParam{Id: data.Id})

	// 验证账号

	if item.Id == 0 {
		return errors.New("账号错误")
	}
	if !user.CheckSecret(item.Password, data.OldPassword) {
		return errors.New("密码错误")
	}
	if err := user.CheckUserinfo(data.Username, data.Password, data.Email); err != nil {
		return err
	}

	// 更新信息

	return user.Update(&data.UpdateParam)

}
