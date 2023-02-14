package config

import (
	"tdp-cloud/module/dborm"
)

// 创建配置

type CreateParam struct {
	Name        string `binding:"required"`
	Value       string `binding:"required"`
	Module      string
	Description string
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Config{
		Name:        data.Name,
		Value:       data.Value,
		Module:      data.Module,
		Description: data.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新配置

type UpdateParam struct {
	Name        string `binding:"required"`
	Value       string `binding:"required"`
	Module      string
	Description string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Config{
			Name: data.Name,
		}).
		Updates(dborm.Config{
			Value:       data.Value,
			Module:      data.Module,
			Description: data.Description,
		})

	return result.Error

}

// 删除配置

type DeleteParam struct {
	Id   uint
	Name string
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&dborm.Config{
			Id:   data.Id,
			Name: data.Name,
		}).
		Delete(&dborm.Config{})

	return result.Error

}

// 获取配置

type FetchParam struct {
	Id   uint
	Name string
}

func Fetch(data *FetchParam) (*dborm.Config, error) {

	var item *dborm.Config

	result := dborm.Db.
		Where(&dborm.Config{
			Id:   data.Id,
			Name: data.Name,
		}).
		First(&item)

	return item, result.Error

}

// 获取配置列表

type FetchAllParam struct {
	Module string
}

func FetchAll(data *FetchAllParam) ([]*dborm.Config, error) {

	var items []*dborm.Config

	result := dborm.Db.
		Where(&dborm.Config{
			Module: data.Module,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.Config{}).
		Where(&dborm.Config{
			Module: data.Module,
		}).
		Count(&count)

	return count, result.Error

}
