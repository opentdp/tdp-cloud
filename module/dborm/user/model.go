package user

import (
	"github.com/google/uuid"

	"tdp-cloud/module/dborm"
)

// 创建用户

type CreateParam struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
	Level    uint
	Email    string `binding:"required"`
}

func Create(data *CreateParam) (uint, error) {

	err := CheckUser(data.Username, data.Password, data.Email)
	if err != nil {
		return 0, err
	}

	item := &dborm.User{
		Username: data.Username,
		Password: HashPassword(data.Password),
		AppId:    uuid.NewString(),
		Level:    data.Level,
		Email:    data.Email,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新用户

type UpdateParam struct {
	Id          uint
	Password    string
	Level       uint
	Email       string
	Description string
}

func Update(data *UpdateParam) error {

	err := CheckUser("", data.Password, data.Email)
	if err != nil {
		return err
	}

	if data.Password != "" {
		data.Password = HashPassword(data.Password)
	}

	result := dborm.Db.
		Where(&dborm.User{
			Id: data.Id,
		}).
		Updates(dborm.User{
			Password:    data.Password,
			Level:       data.Level,
			Email:       data.Email,
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
	AppId    string
	Username string
	Email    string
}

func Fetch(data *FetchParam) (*dborm.User, error) {

	var item *dborm.User

	result := dborm.Db.
		Where(&dborm.User{
			Id:       data.Id,
			AppId:    data.AppId,
			Username: data.Username,
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
