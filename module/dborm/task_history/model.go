package task_history

import (
	"tdp-cloud/module/dborm"
)

// 添加任务

type CreateParam struct {
	UserId   uint
	HostId   string `binding:"required"`
	HostName string `binding:"required"`
	Subject  string `binding:"required"`
	Status   string `binding:"required"`
	Request  any    `binding:"required"`
	Response any
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.TaskHistory{
		UserId:   post.UserId,
		HostId:   post.HostId,
		HostName: post.HostName,
		Subject:  post.Subject,
		Status:   post.Status,
		Request:  post.Request,
		Response: post.Response,
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
	Request  any
	Response any
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.TaskHistory{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.TaskHistory{
			HostId:   post.HostId,
			HostName: post.HostName,
			Subject:  post.Subject,
			Status:   post.Status,
			Request:  post.Request,
			Response: post.Response,
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
