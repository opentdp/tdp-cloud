package domain

import (
	"tdp-cloud/module/dborm"
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
	Description string
	Status      uint
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Domain{
		UserId:      data.UserId,
		VendorId:    data.VendorId,
		Name:        data.Name,
		NSList:      data.NSList,
		Model:       data.Model,
		CloudId:     data.CloudId,
		CloudMeta:   data.CloudMeta,
		Description: data.Description,
		Status:      data.Status,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新域名

type UpdateParam struct {
	Id          uint `binding:"required"`
	UserId      uint
	VendorId    uint
	Name        string
	NSList      string
	Model       string
	CloudId     string
	CloudMeta   any
	Description string
	Status      uint
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Domain{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(dborm.Domain{
			VendorId:    data.VendorId,
			Name:        data.Name,
			NSList:      data.NSList,
			Model:       data.Model,
			CloudId:     data.CloudId,
			CloudMeta:   data.CloudMeta,
			Description: data.Description,
			Status:      data.Status,
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
		Where(&dborm.Domain{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&dborm.Domain{})

	return result.Error

}

// 获取域名

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*dborm.Domain, error) {

	var item *dborm.Domain

	result := dborm.Db.
		Where(&dborm.Domain{
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

func FetchAll(data *FetchAllParam) ([]*dborm.Domain, error) {

	var items []*dborm.Domain

	result := dborm.Db.
		Where(&dborm.Domain{
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
		Where(&dborm.Domain{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Count(&count)

	return count, result.Error

}
