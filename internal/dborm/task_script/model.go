package task_script

import (
	"tdp-cloud/internal/dborm"
)

// 添加脚本

type CreateParam struct {
	UserId        uint
	Name          string `binding:"required"`
	Username      string `binding:"required"`
	Description   string
	Content       string `binding:"required"`
	CommandType   string `binding:"required"`
	WorkDirectory string `binding:"required"`
	Timeout       uint   `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.TaskScript{
		UserId:        post.UserId,
		Name:          post.Name,
		Username:      post.Username,
		Description:   post.Description,
		Content:       post.Content,
		WorkDirectory: post.WorkDirectory,
		CommandType:   post.CommandType,
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
	Username      string `binding:"required"`
	Description   string
	Content       string `binding:"required"`
	CommandType   string `binding:"required"`
	WorkDirectory string `binding:"required"`
	Timeout       uint   `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.TaskScript{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.TaskScript{
			Name:          post.Name,
			Username:      post.Username,
			Description:   post.Description,
			Content:       post.Content,
			CommandType:   post.CommandType,
			WorkDirectory: post.WorkDirectory,
			Timeout:       post.Timeout,
		})

	return result.Error

}

// 获取脚本列表

func FetchAll(userId uint) ([]*dborm.TaskScript, error) {

	var items []*dborm.TaskScript

	result := dborm.Db.Where("user_id = 0").Or(&dborm.TaskScript{UserId: userId}).Find(&items)

	return items, result.Error

}

// 获取脚本

func Fetch(id, userId uint) (*dborm.TaskScript, error) {

	var item *dborm.TaskScript

	result := dborm.Db.Where(&dborm.TaskScript{Id: id, UserId: userId}).Find(&item)

	return item, result.Error

}

// 删除脚本

func Delete(id, userId uint) error {

	var item *dborm.TaskScript

	result := dborm.Db.Where(&dborm.TaskScript{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
