package domain

import (
	"gorm.io/datatypes"

	"tdp-cloud/module/dborm"
)

// 添加域名

type CreateParam struct {
	UserId      uint
	VendorId    uint   `binding:"required"`
	Name        string `binding:"required"`
	NSList      string `binding:"required"`
	Model       string `binding:"required"`
	CloudId     string
	CloudMeta   string
	Description string
	Status      string `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Domain{
		UserId:      post.UserId,
		VendorId:    post.VendorId,
		Name:        post.Name,
		NSList:      post.NSList,
		Model:       post.Model,
		CloudId:     post.CloudId,
		CloudMeta:   datatypes.JSON(post.CloudMeta),
		Description: post.Description,
		Status:      datatypes.JSON(post.Status),
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新域名

type UpdateParam struct {
	Id          uint `binding:"required"`
	UserId      uint
	VendorId    uint   `binding:"required"`
	Name        string `binding:"required"`
	NSList      string `binding:"required"`
	Model       string `binding:"required"`
	CloudId     string
	CloudMeta   string
	Description string
	Status      string `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Domain{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Domain{
			VendorId:    post.VendorId,
			Name:        post.Name,
			NSList:      post.NSList,
			Model:       post.Model,
			CloudId:     post.CloudId,
			CloudMeta:   datatypes.JSON(post.CloudMeta),
			Description: post.Description,
			Status:      datatypes.JSON(post.Status),
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
