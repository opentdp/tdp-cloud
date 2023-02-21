package cronjob

import (
	"tdp-cloud/module/dborm"
)

// 创建计划

type CreateParam struct {
	UserId     uint
	Name       string `binding:"required"`
	Type       string `binding:"required"`
	Content    string `binding:"required"`
	Second     string `binding:"required"`
	Minute     string `binding:"required"`
	Hour       string `binding:"required"`
	DayofMonth string `binding:"required"`
	Month      string `binding:"required"`
	DayofWeek  string `binding:"required"`
	Location   string `binding:"required"`
	PrevTime   int64  `binding:"required"`
	NextTime   int64  `binding:"required"`
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Cronjob{
		UserId:     data.UserId,
		Name:       data.Name,
		Type:       data.Type,
		Content:    data.Content,
		Second:     data.Second,
		Minute:     data.Minute,
		Hour:       data.Hour,
		DayofMonth: data.DayofMonth,
		Month:      data.Month,
		DayofWeek:  data.DayofWeek,
		Location:   data.Location,
		PrevTime:   data.PrevTime,
		NextTime:   data.NextTime,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新计划

type UpdateParam struct {
	Id         uint `binding:"required"`
	UserId     uint
	Name       string
	Type       string
	Content    string
	Second     string
	Minute     string
	Hour       string
	DayofMonth string
	Month      string
	DayofWeek  string
	Location   string
	PrevTime   int64
	NextTime   int64
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Cronjob{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(dborm.Cronjob{
			Name:       data.Name,
			Type:       data.Type,
			Content:    data.Content,
			Second:     data.Second,
			Minute:     data.Minute,
			Hour:       data.Hour,
			DayofMonth: data.DayofMonth,
			Month:      data.Month,
			DayofWeek:  data.DayofWeek,
			Location:   data.Location,
			PrevTime:   data.PrevTime,
			NextTime:   data.NextTime,
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
		Where(&dborm.Cronjob{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&dborm.Cronjob{})

	return result.Error

}

// 获取计划

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*dborm.Cronjob, error) {

	var item *dborm.Cronjob

	result := dborm.Db.
		Where(&dborm.Cronjob{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取计划列表

type FetchAllParam struct {
	UserId   uint
	VendorId uint
}

func FetchAll(data *FetchAllParam) ([]*dborm.Cronjob, error) {

	var items []*dborm.Cronjob

	result := dborm.Db.
		Where(&dborm.Cronjob{
			UserId: data.UserId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取计划总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.Cronjob{}).
		Where(&dborm.Cronjob{
			UserId: data.UserId,
		}).
		Count(&count)

	return count, result.Error

}
