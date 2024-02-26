package config

import (
	"github.com/opentdp/go-helper/dborm"

	"tdp-cloud/model"
)

// 创建配置

type CreateParam struct {
	Name        string `binding:"required"`
	Value       string `binding:"required"`
	Type        string
	Module      string
	Description string
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Config{
		Name:        data.Name,
		Value:       data.Value,
		Type:        data.Type,
		Module:      data.Module,
		Description: data.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新配置

type UpdateParam struct {
	Id          uint
	Name        string
	Value       string
	Type        string
	Module      string
	Description string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Config{
			Id: data.Id,
		}).
		Updates(model.Config{
			Name:        data.Name,
			Value:       data.Value,
			Type:        data.Type,
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
		Where(&model.Config{
			Id:   data.Id,
			Name: data.Name,
		}).
		Delete(&model.Config{})

	return result.Error

}

// 获取配置

type FetchParam struct {
	Id   uint
	Name string
}

func Fetch(data *FetchParam) (*model.Config, error) {

	var item *model.Config

	result := dborm.Db.
		Where(&model.Config{
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

func FetchAll(data *FetchAllParam) ([]*model.Config, error) {

	var items []*model.Config

	result := dborm.Db.
		Where(&model.Config{
			Module: data.Module,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.Config{}).
		Where(&model.Config{
			Module: data.Module,
		}).
		Count(&count)

	return count, result.Error

}
