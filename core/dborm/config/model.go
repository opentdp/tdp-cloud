package config

import (
	"tdp-cloud/core/dborm"
)

// 添加配置

type CreateParam struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func Create(post *CreateParam) error {

	result := dborm.Db.Create(&dborm.Config{
		Key:   post.Key,
		Value: post.Value,
	})

	return result.Error

}

// 更新配置

type UpdateParam struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.Model(&dborm.Config{}).
		Where("key = ? ", post.Key).
		Updates(dborm.Config{
			Value: post.Value,
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

func Fetch(key string) (*dborm.Config, error) {

	var item *dborm.Config

	result := dborm.Db.First(&item, "key = ?", key)

	return item, result.Error

}

// 删除配置

func Delete(key string) error {

	var item *dborm.Config

	result := dborm.Db.Delete(&item, "key = ?", key)

	return result.Error

}
