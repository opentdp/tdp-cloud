package vendor

import (
	"tdp-cloud/internal/dborm"
)

// 添加厂商

type CreateParam struct {
	UserId      uint
	SecretId    string `binding:"required"`
	SecretKey   string `binding:"required"`
	Provider    string `binding:"required"`
	Description string `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Vendor{
		UserId:      post.UserId,
		SecretId:    post.SecretId,
		SecretKey:   post.SecretKey,
		Provider:    post.Provider,
		Description: post.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新厂商

type UpdateParam struct {
	Id          uint `binding:"required"`
	UserId      uint
	SecretId    string `binding:"required"`
	SecretKey   string `binding:"required"`
	Provider    string `binding:"required"`
	Description string `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Vendor{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Vendor{
			SecretId:    post.SecretId,
			SecretKey:   post.SecretKey,
			Provider:    post.Provider,
			Description: post.Description,
		})

	return result.Error

}

// 获取厂商列表

func FetchAll(userId uint) ([]*dborm.Vendor, error) {

	var items []*dborm.Vendor

	result := dborm.Db.Where(&dborm.Vendor{UserId: userId}).Find(&items)

	return items, result.Error

}

// 获取厂商

func Fetch(id, userId uint) (*dborm.Vendor, error) {

	var item *dborm.Vendor

	result := dborm.Db.Where(&dborm.Vendor{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除厂商

func Delete(id, userId uint) error {

	var item *dborm.Vendor

	result := dborm.Db.Where(&dborm.Vendor{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
