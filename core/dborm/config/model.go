package config

import (
	"tdp-cloud/core/dborm"
)

// 添加配置

type CreateParam struct {
	Name        string `binding:"required"`
	Value       string `binding:"required"`
	Module      string
	Description string
}

func Create(post *CreateParam) error {

	result := dborm.Db.Create(&dborm.Config{
		Name:        post.Name,
		Value:       post.Value,
		Module:      post.Module,
		Description: post.Description,
	})

	return result.Error

}

// 更新配置

type UpdateParam struct {
	Name        string `binding:"required"`
	Value       string `binding:"required"`
	Module      string
	Description string
}

func Update(post *UpdateParam) error {

	result := dborm.Db.Model(&dborm.Config{}).
		Where("name = ? ", post.Name).
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

	result := dborm.Db.First(&item, "name = ?", name)

	return item, result.Error

}

// 删除配置

func Delete(name string) error {

	var item *dborm.Config

	result := dborm.Db.Delete(&item, "name = ?", name)

	return result.Error

}
