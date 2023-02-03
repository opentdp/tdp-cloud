package script

import (
	"tdp-cloud/module/dborm"
)

// 添加脚本

type CreateParam struct {
	UserId        uint
	Name          string `binding:"required"`
	CommandType   string `binding:"required"`
	Username      string `binding:"required"`
	WorkDirectory string `binding:"required"`
	Content       string `binding:"required"`
	Description   string
	Timeout       uint `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Script{
		UserId:        post.UserId,
		Name:          post.Name,
		CommandType:   post.CommandType,
		Username:      post.Username,
		WorkDirectory: post.WorkDirectory,
		Content:       post.Content,
		Description:   post.Description,
		Timeout:       post.Timeout,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新脚本

type UpdateParam struct {
	Id            uint `binding:"required"`
	UserId        uint
	Name          string `binding:"required"`
	CommandType   string `binding:"required"`
	Username      string `binding:"required"`
	WorkDirectory string `binding:"required"`
	Content       string `binding:"required"`
	Description   string
	Timeout       uint `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Script{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Script{
			Name:          post.Name,
			CommandType:   post.CommandType,
			Username:      post.Username,
			WorkDirectory: post.WorkDirectory,
			Content:       post.Content,
			Description:   post.Description,
			Timeout:       post.Timeout,
		})

	return result.Error

}

// 获取脚本列表

func FetchAll(userId uint) ([]*dborm.Script, error) {

	var items []*dborm.Script

	result := dborm.Db.Where(&dborm.Script{UserId: userId}).Find(&items)

	return items, result.Error

}

// 获取脚本

func Fetch(id, userId uint) (*dborm.Script, error) {

	var item *dborm.Script

	result := dborm.Db.Where(&dborm.Script{Id: id, UserId: userId}).Find(&item)

	return item, result.Error

}

// 删除脚本

func Delete(id, userId uint) error {

	var item *dborm.Script

	result := dborm.Db.Where(&dborm.Script{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
