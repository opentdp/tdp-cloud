package slave_task

import (
	"tdp-cloud/core/dborm"
)

// 添加任务

type CreateParam struct {
	UserId   uint
	HostId   string `binding:"required"`
	HostName string `binding:"required"`
	Subject  string `binding:"required"`
	Content  string `binding:"required"`
	Status   string `binding:"required"`
	Result   string
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.SlaveTask{
		UserId:   post.UserId,
		HostId:   post.HostId,
		HostName: post.HostName,
		Subject:  post.Subject,
		Content:  post.Content,
		Status:   post.Status,
		Result:   post.Result,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新任务

type UpdateParam struct {
	Id       uint `binding:"required"`
	UserId   uint
	HostId   string `binding:"required"`
	HostName string `binding:"required"`
	Subject  string `binding:"required"`
	Content  string `binding:"required"`
	Status   string `binding:"required"`
	Result   string
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.SlaveTask{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.SlaveTask{
			HostId:   post.HostId,
			HostName: post.HostName,
			Subject:  post.Subject,
			Content:  post.Content,
			Status:   post.Status,
			Result:   post.Result,
		})

	return result.Error

}

// 获取任务列表

func FetchAll(userId uint) ([]*dborm.SlaveTask, error) {

	var items []*dborm.SlaveTask

	result := dborm.Db.
		Where(&dborm.SlaveTask{UserId: userId}).
		Limit(50).Order("id DESC").
		Find(&items)

	return items, result.Error

}

// 获取任务

func Fetch(id, userId uint) (*dborm.SlaveTask, error) {

	var item *dborm.SlaveTask

	result := dborm.Db.Where(&dborm.SlaveTask{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除任务

func Delete(id, userId uint) error {

	var item *dborm.SlaveTask

	result := dborm.Db.Where(&dborm.SlaveTask{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
