package taskline

import (
	"tdp-cloud/module/dborm"
)

// 创建任务

type CreateParam struct {
	UserId   uint
	Subject  string `binding:"required"`
	HostName string `binding:"required"`
	WorkerId string `binding:"required"`
	Status   string `binding:"required"`
	Request  any    `binding:"required"`
	Response any
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Taskline{
		UserId:   data.UserId,
		Subject:  data.Subject,
		HostName: data.HostName,
		WorkerId: data.WorkerId,
		Status:   data.Status,
		Request:  data.Request,
		Response: data.Response,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新任务

type UpdateParam struct {
	Id       uint
	UserId   uint
	Subject  string
	HostName string
	WorkerId string
	Status   string
	Request  any
	Response any
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Taskline{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(dborm.Taskline{
			Subject:  data.Subject,
			HostName: data.HostName,
			WorkerId: data.WorkerId,
			Status:   data.Status,
			Request:  data.Request,
			Response: data.Response,
		})

	return result.Error

}

// 删除任务

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	var item *dborm.Taskline

	result := dborm.Db.
		Where(&dborm.Taskline{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&item)

	return result.Error

}

// 获取任务

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*dborm.Taskline, error) {

	var item *dborm.Taskline

	result := dborm.Db.
		Where(&dborm.Taskline{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取任务列表

type FetchAllParam struct {
	UserId   uint
	WorkerId string
}

func FetchAll(data *FetchAllParam) ([]*dborm.Taskline, error) {

	var items []*dborm.Taskline

	result := dborm.Db.
		Where(&dborm.Taskline{
			UserId:   data.UserId,
			WorkerId: data.WorkerId,
		}).
		Order("id DESC").
		Limit(50).
		Find(&items)

	return items, result.Error

}

// 获取任务总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.Taskline{}).
		Where(&dborm.Taskline{
			UserId:   data.UserId,
			WorkerId: data.WorkerId,
		}).
		Count(&count)

	return count, result.Error

}
