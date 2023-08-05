package taskline

import (
	"github.com/opentdp/go-helper/dborm"

	"tdp-cloud/model"
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

	item := &model.Taskline{
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
		Where(&model.Taskline{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(model.Taskline{
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

	var item *model.Taskline

	result := dborm.Db.
		Where(&model.Taskline{
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

func Fetch(data *FetchParam) (*model.Taskline, error) {

	var item *model.Taskline

	result := dborm.Db.
		Where(&model.Taskline{
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
	Page     int
	Order    string
}

func FetchAll(data *FetchAllParam) ([]*model.Taskline, error) {

	var items []*model.Taskline

	offset := 0
	if data.Page > 1 {
		offset = (data.Page - 1) * 100
	}

	if err := dborm.OrderSafe(data.Order); err != nil {
		return nil, err
	}

	result := dborm.Db.
		Where(&model.Taskline{
			UserId:   data.UserId,
			WorkerId: data.WorkerId,
		}).
		Order(data.Order).
		Limit(100).Offset(offset).
		Find(&items)

	return items, result.Error

}

// 获取任务总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.Taskline{}).
		Where(&model.Taskline{
			UserId:   data.UserId,
			WorkerId: data.WorkerId,
		}).
		Count(&count)

	return count, result.Error

}
