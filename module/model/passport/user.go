package passport

import (
	"errors"

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

	if item.Id == 0 {
		return nil, errors.New("账号错误")
	}
	if !user.CheckPassword(item.Password, data.Password) {
		return nil, errors.New("密码错误")
	}

	// 创建令牌

	token, _ := midware.CreateToken(&midware.UserInfo{
		UserId:    item.Id,
		UserLevel: item.Level,
		SecretKey: "",
	})

	// 返回结果

	res := &LoginResult{
		Username: item.Username,
		AppId:    item.AppId,
		Email:    item.Email,
		Token:    token,
	}

	return res, nil

}

// 修改密码

type UpdateInfoParam struct {
	Id          uint
	Username    string `binding:"required"`
	Password    string
	Email       string `binding:"required"`
	Description string
	OldPassword string `binding:"required"`
}

func UpdateInfo(data *UpdateInfoParam) error {

	item, _ := user.Fetch(&user.FetchParam{Id: data.Id})

	// 验证账号

	if item.Id == 0 {
		return errors.New("账号错误")
	}
	if !user.CheckPassword(item.Password, data.OldPassword) {
		return errors.New("密码错误")
	}
	if err := user.CheckUser(data.Username, data.Password, data.Email); err != nil {
		return err
	}

	// 更新信息

	return user.Update(&user.UpdateParam{
		Id:          data.Id,
		Username:    data.Username,
		Password:    data.Password,
		Email:       data.Email,
		Description: data.Description,
	})

}
