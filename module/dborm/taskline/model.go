package taskline

import (
	"tdp-cloud/module/dborm"
)

// 添加任务

type CreateParam struct {
	UserId   uint
	Subject  string `binding:"required"`
	HostName string `binding:"required"`
	WorkerId string `binding:"required"`
	Status   string `binding:"required"`
	Request  any    `binding:"required"`
	Response any
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Taskline{
		UserId:   post.UserId,
		Subject:  post.Subject,
		HostName: post.HostName,
		WorkerId: post.WorkerId,
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
	Subject  string
	HostName string
	WorkerId string
	Status   string `binding:"required"`
	Request  any
	Response any
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Taskline{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Taskline{
			Subject:  post.Subject,
			HostName: post.HostName,
			WorkerId: post.WorkerId,
			Status:   post.Status,
			Request:  post.Request,
			Response: post.Response,
		})

	return result.Error

}

// 获取任务列表

func FetchAll(userId uint) ([]*dborm.Taskline, error) {

	var items []*dborm.Taskline

	result := dborm.Db.
		Where(&dborm.Taskline{UserId: userId}).
		Limit(50).Order("id DESC").
		Find(&items)

	return items, result.Error

}

// 获取任务

func Fetch(id, userId uint) (*dborm.Taskline, error) {

	var item *dborm.Taskline

	result := dborm.Db.Where(&dborm.Taskline{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除任务

func Delete(id, userId uint) error {

	var item *dborm.Taskline

	result := dborm.Db.Where(&dborm.Taskline{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
