package user

import (
	"github.com/google/uuid"

	"tdp-cloud/helper/strutil"
	"tdp-cloud/module/dborm"
)

// 创建用户

type CreateParam struct {
	Username    string `binding:"required"`
	Password    string `binding:"required"`
	Level       uint
	AppKey      string
	Email       string `binding:"required"`
	Description string
}

func Create(data *CreateParam) (uint, error) {

	if len(data.AppKey) != 32 {
		data.AppKey = strutil.Rand(32) //强制为32位
	}

	sk, pw, err := NewSecret(data.AppKey, data.Password)
	if err != nil {
		return 0, err
	}

	item := &dborm.User{
		Username:    data.Username,
		Password:    pw,
		Level:       data.Level,
		AppId:       uuid.NewString(),
		AppKey:      sk,
		Email:       data.Email,
		Description: data.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新用户

type UpdateParam struct {
	Id          uint
	Username    string
	Password    string
	Level       uint
	Email       string
	AppKey      string
	Description string
}

func Update(data *UpdateParam) error {

	sk, pw, err := NewSecret(data.AppKey, data.Password)
	if err != nil {
		return err
	}

	result := dborm.Db.
		Where(&dborm.User{
			Id: data.Id,
		}).
		Updates(dborm.User{
			Username:    data.Username,
			Password:    pw,
			Level:       data.Level,
			Email:       data.Email,
			AppKey:      sk,
			Description: data.Description,
		})

	return result.Error

}

// 删除用户

type DeleteParam struct {
	Id       uint
	Username string
}

func Delete(data *DeleteParam) error {

	var item *dborm.Taskline

	result := dborm.Db.
		Where(&dborm.User{
			Id:       data.Id,
			Username: data.Username,
		}).
		Delete(&item)

	return result.Error

}

// 获取用户

type FetchParam struct {
	Id       uint
	Username string
	AppId    string
	Email    string
}

func Fetch(data *FetchParam) (*dborm.User, error) {

	var item *dborm.User

	result := dborm.Db.
		Where(&dborm.User{
			Id:       data.Id,
			Username: data.Username,
			AppId:    data.AppId,
			Email:    data.Email,
		}).
		First(&item)

	return item, result.Error

}

// 获取用户列表

type FetchAllParam struct {
	Level uint
}

func FetchAll(data *FetchAllParam) ([]*dborm.User, error) {

	var items []*dborm.User

	result := dborm.Db.
		Where(&dborm.User{
			Level: data.Level,
		}).
		Find(&items)

	return items, result.Error

}

// 获取用户总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.User{}).
		Where(&dborm.User{
			Level: data.Level,
		}).
		Count(&count)

	return count, result.Error

}
