package user

import (
	"github.com/google/uuid"

	"tdp-cloud/module/dborm"
)

// 添加用户

type CreateParam struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	if post.Password != "" {
		post.Password = HashPassword(post.Password)
	}

	item := &dborm.User{
		AppId:    uuid.NewString(),
		Username: post.Username,
		Password: post.Password,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新用户

type UpdateParam struct {
	Id          uint
	Description string
	Password    string
}

func Update(post *UpdateParam) error {

	if post.Password != "" {
		post.Password = HashPassword(post.Password)
	}

	result := dborm.Db.
		Where(&dborm.User{Id: post.Id}).
		Updates(dborm.User{
			Description: post.Description,
			Password:    post.Password,
		})

	return result.Error

}

// 获取用户列表

func FetchAll() ([]*dborm.User, error) {

	var items []*dborm.User

	result := dborm.Db.Find(&items)

	return items, result.Error

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

	return item, result.Error

}
