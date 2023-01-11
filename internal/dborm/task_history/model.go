package task_history

import (
	"tdp-cloud/internal/dborm"

	"gorm.io/datatypes"
)

// 添加任务

type CreateParam struct {
	UserId   uint
	HostId   string `binding:"required"`
	HostName string `binding:"required"`
	Subject  string `binding:"required"`
	Status   string `binding:"required"`
	Request  string `binding:"required"`
	Response string
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.TaskHistory{
		UserId:   post.UserId,
		HostId:   post.HostId,
		HostName: post.HostName,
		Subject:  post.Subject,
		Status:   post.Status,
		Request:  datatypes.JSON(post.Request),
		Response: datatypes.JSON(post.Response),
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新任务

type UpdateParam struct {
	Id       uint `binding:"required"`
	UserId   uint
	HostId   string
	HostName string
	Subject  string
	Status   string `binding:"required"`
	Request  string
	Response string
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.TaskHistory{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.TaskHistory{
			HostId:   post.HostId,
			HostName: post.HostName,
			Subject:  post.Subject,
			Status:   post.Status,
			Request:  datatypes.JSON(post.Request),
			Response: datatypes.JSON(post.Response),
		})

	return result.Error

}

// 获取任务列表

func FetchAll(userId uint) ([]*dborm.TaskHistory, error) {

	var items []*dborm.TaskHistory

	result := dborm.Db.
		Where(&dborm.TaskHistory{UserId: userId}).
		Limit(50).Order("id DESC").
		Find(&items)

	return items, result.Error

}

// 获取任务

func Fetch(id, userId uint) (*dborm.TaskHistory, error) {

	var item *dborm.TaskHistory

	result := dborm.Db.Where(&dborm.TaskHistory{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除任务

func Delete(id, userId uint) error {

	var item *dborm.TaskHistory

	result := dborm.Db.Where(&dborm.TaskHistory{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
