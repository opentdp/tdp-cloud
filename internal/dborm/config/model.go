package config

import (
	"tdp-cloud/internal/dborm"
)

// 添加配置

type CreateParam struct {
	Name        string `binding:"required"`
	Value       string `binding:"required"`
	Module      string
	Description string
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Config{
		Name:        post.Name,
		Value:       post.Value,
		Module:      post.Module,
		Description: post.Description,
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

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Config{Name: post.Name}).
		Updates(dborm.Config{
			Value:       post.Value,
			Module:      post.Module,
			Description: post.Description,
		})

	return result.Error

}

// 获取配置列表

func FetchAll() ([]*dborm.Config, error) {

	var items []*dborm.Config

	result := dborm.Db.Find(&items)

	return items, result.Error

}

// 获取配置

func Fetch(name string) (*dborm.Config, error) {

	var item *dborm.Config

	result := dborm.Db.Where(&dborm.Config{Name: name}).First(&item)

	return item, result.Error

}

// 删除配置

func Delete(name string) error {

	var item *dborm.Config

	result := dborm.Db.Where(&dborm.Config{Name: name}).Delete(&item)

	return result.Error

}
