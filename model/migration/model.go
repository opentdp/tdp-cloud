package migration

import (
	"github.com/opentdp/go-helper/dborm"

	"tdp-cloud/model"
)

// 创建配置

type CreateParam struct {
	Version     string `binding:"required"`
	Description string
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Migration{
		Version:     data.Version,
		Description: data.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新配置

type UpdateParam struct {
	Id          uint
	Version     string
	Description string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Migration{
			Id: data.Id,
		}).
		Updates(model.Migration{
			Version:     data.Version,
			Description: data.Description,
		})

	return result.Error

}

// 删除配置

type DeleteParam struct {
	Id      uint
	Version string
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&model.Migration{
			Id:      data.Id,
			Version: data.Version,
		}).
		Delete(&model.Migration{})

	return result.Error

}

// 获取配置

type FetchParam struct {
	Id      uint
	Version string
}

func Fetch(data *FetchParam) (*model.Migration, error) {

	var item *model.Migration

	result := dborm.Db.
		Where(&model.Migration{
			Id:      data.Id,
			Version: data.Version,
		}).
		First(&item)

	return item, result.Error

}
