package cronjob

import (
	"github.com/opentdp/go-helper/dborm"

	"tdp-cloud/model"
)

// 创建计划

type CreateParam struct {
	UserId     uint
	Name       string `binding:"required"`
	Type       string `binding:"required"`
	Target     string `binding:"required"`
	Content    string `binding:"required"`
	Second     string `binding:"required"`
	Minute     string `binding:"required"`
	Hour       string `binding:"required"`
	DayofMonth string `binding:"required"`
	Month      string `binding:"required"`
	DayofWeek  string `binding:"required"`
	Location   string
	EntryId    int64
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Cronjob{
		UserId:     data.UserId,
		Name:       data.Name,
		Type:       data.Type,
		Target:     data.Target,
		Content:    data.Content,
		Second:     data.Second,
		Minute:     data.Minute,
		Hour:       data.Hour,
		DayofMonth: data.DayofMonth,
		Month:      data.Month,
		DayofWeek:  data.DayofWeek,
		Location:   data.Location,
		EntryId:    data.EntryId,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新计划

type UpdateParam struct {
	Id         uint
	UserId     uint
	Name       string
	Type       string
	Target     string
	Content    string
	Second     string
	Minute     string
	Hour       string
	DayofMonth string
	Month      string
	DayofWeek  string
	Location   string
	EntryId    int64
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Cronjob{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(model.Cronjob{
			Name:       data.Name,
			Type:       data.Type,
			Target:     data.Target,
			Content:    data.Content,
			Second:     data.Second,
			Minute:     data.Minute,
			Hour:       data.Hour,
			DayofMonth: data.DayofMonth,
			Month:      data.Month,
			DayofWeek:  data.DayofWeek,
			Location:   data.Location,
			EntryId:    data.EntryId,
		})

	return result.Error

}

// 删除计划

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&model.Cronjob{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&model.Cronjob{})

	return result.Error

}

// 获取计划

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*model.Cronjob, error) {

	var item *model.Cronjob

	result := dborm.Db.
		Where(&model.Cronjob{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取计划列表

type FetchAllParam struct {
	UserId uint
}

func FetchAll(data *FetchAllParam) ([]*model.Cronjob, error) {

	var items []*model.Cronjob

	result := dborm.Db.
		Where(&model.Cronjob{
			UserId: data.UserId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取计划总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.Cronjob{}).
		Where(&model.Cronjob{
			UserId: data.UserId,
		}).
		Count(&count)

	return count, result.Error

}
