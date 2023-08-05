package domain

import (
	"github.com/opentdp/go-helper/dborm"

	"tdp-cloud/model"
)

// 创建域名

type CreateParam struct {
	UserId      uint
	VendorId    uint   `binding:"required"`
	Name        string `binding:"required"`
	NSList      string `binding:"required"`
	Model       string `binding:"required"`
	CloudId     string
	CloudMeta   any
	Status      string
	Description string
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Domain{
		UserId:      data.UserId,
		VendorId:    data.VendorId,
		Name:        data.Name,
		NSList:      data.NSList,
		Model:       data.Model,
		CloudId:     data.CloudId,
		CloudMeta:   data.CloudMeta,
		Status:      data.Status,
		Description: data.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新域名

type UpdateParam struct {
	Id          uint
	UserId      uint
	VendorId    uint
	Name        string
	NSList      string
	Model       string
	CloudId     string
	CloudMeta   any
	Status      string
	Description string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Domain{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(model.Domain{
			VendorId:    data.VendorId,
			Name:        data.Name,
			NSList:      data.NSList,
			Model:       data.Model,
			CloudId:     data.CloudId,
			CloudMeta:   data.CloudMeta,
			Status:      data.Status,
			Description: data.Description,
		})

	return result.Error

}

// 删除域名

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&model.Domain{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&model.Domain{})

	return result.Error

}

// 获取域名

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*model.Domain, error) {

	var item *model.Domain

	result := dborm.Db.
		Where(&model.Domain{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取域名列表

type FetchAllParam struct {
	UserId   uint
	VendorId uint
}

func FetchAll(data *FetchAllParam) ([]*model.Domain, error) {

	var items []*model.Domain

	result := dborm.Db.
		Where(&model.Domain{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取域名总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.Domain{}).
		Where(&model.Domain{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Count(&count)

	return count, result.Error

}
