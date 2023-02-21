package certbot

import (
	"tdp-cloud/module/dborm"
)

// 创建计划

type CreateParam struct {
	UserId   uint
	VendorId uint   `binding:"required"`
	Email    string `binding:"required"`
	Domain   string `binding:"required"`
}

func Create(data *CreateParam) (uint, error) {

	item := &dborm.Certbot{
		UserId:   data.UserId,
		VendorId: data.VendorId,
		Email:    data.Email,
		Domain:   data.Domain,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新计划

type UpdateParam struct {
	Id       uint `binding:"required"`
	UserId   uint
	VendorId uint
	Email    string
	Domain   string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Certbot{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(dborm.Certbot{
			VendorId: data.VendorId,
			Email:    data.Email,
			Domain:   data.Domain,
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
		Where(&dborm.Certbot{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&dborm.Certbot{})

	return result.Error

}

// 获取计划

type FetchParam struct {
	Id     uint
	UserId uint
}

func Fetch(data *FetchParam) (*dborm.Certbot, error) {

	var item *dborm.Certbot

	result := dborm.Db.
		Where(&dborm.Certbot{
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

func FetchAll(data *FetchAllParam) ([]*dborm.Certbot, error) {

	var items []*dborm.Certbot

	result := dborm.Db.
		Where(&dborm.Certbot{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取计划总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&dborm.Certbot{}).
		Where(&dborm.Certbot{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Count(&count)

	return count, result.Error

}
