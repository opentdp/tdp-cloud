package domain

import (
	"tdp-cloud/internal/dborm"
)

// 添加域名

type CreateParam struct {
	UserId      uint   `binding:"required"`
	VendorId    uint   `binding:"required"`
	Name        string `binding:"required"`
	Status      string
	CloudData   string `binding:"required"`
	Description string `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Domain{
		UserId:      post.UserId,
		VendorId:    post.VendorId,
		Name:        post.Name,
		Status:      post.Status,
		CloudData:   post.CloudData,
		Description: post.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新域名

type UpdateParam struct {
	Id          uint   `binding:"required"`
	UserId      uint   `binding:"required"`
	VendorId    uint   `binding:"required"`
	Name        string `binding:"required"`
	Status      string
	CloudData   string `binding:"required"`
	Description string `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Domain{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Domain{
			VendorId:    post.VendorId,
			Name:        post.Name,
			Status:      post.Status,
			CloudData:   post.CloudData,
			Description: post.Description,
		})

	return result.Error

}

// 获取域名列表

func FetchAll(userId uint) ([]*dborm.Domain, error) {

	var items []*dborm.Domain

	result := dborm.Db.Where(&dborm.Domain{UserId: userId}).Find(&items)

	return items, result.Error

}

// 获取域名

func Fetch(id, userId uint) (*dborm.Domain, error) {

	var item *dborm.Domain

	result := dborm.Db.Where(&dborm.Domain{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除域名

func Delete(id, userId uint) error {

	var item *dborm.Domain

	result := dborm.Db.Where(&dborm.Domain{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
