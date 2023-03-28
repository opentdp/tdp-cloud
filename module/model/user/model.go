package user

import (
	"github.com/google/uuid"

	"tdp-cloud/helper/secure"
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
	StoreKey    string // 存储密钥
}

func Create(data *CreateParam) (uint, error) {

	if data.Password != "" {
		pw, err := CreateSecret(data.Password)
		if err != nil {
			return 0, err
		}
		data.Password = pw
	}

	if data.AppKey != "" && data.StoreKey != "" {
		secret, err := secure.Des3Encrypt(data.AppKey, data.StoreKey)
		if err == nil {
			data.AppKey = secret
		}
	}

	item := &dborm.User{
		Username:    data.Username,
		Password:    data.Password,
		Level:       data.Level,
		AppId:       uuid.NewString(),
		AppKey:      data.AppKey,
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
	StoreKey    string // 存储密钥
}

func Update(data *UpdateParam) error {

	if data.Password != "" {
		pw, err := CreateSecret(data.Password)
		if err != nil {
			return err
		}
		data.Password = pw
	}

	if data.AppKey != "" && data.StoreKey != "" {
		secret, err := secure.Des3Encrypt(data.AppKey, data.StoreKey)
		if err == nil {
			data.AppKey = secret
		}
	}

	result := dborm.Db.
		Where(&dborm.User{
			Id: data.Id,
		}).
		Updates(dborm.User{
			Username:    data.Username,
			Password:    data.Password,
			Level:       data.Level,
			Email:       data.Email,
			AppKey:      data.AppKey,
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

	var item *dborm.User

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
	StoreKey string // 存储密钥
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

	if item.AppKey != "" && data.StoreKey != "" {
		item.AppKey, _ = secure.Des3Decrypt(item.AppKey, data.StoreKey)
	}

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
