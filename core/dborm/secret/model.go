package secret

import (
	"tdp-cloud/core/dborm"
)

// 添加密钥

type CreateParam struct {
	UserId      uint
	SecretId    string `binding:"required"`
	SecretKey   string `binding:"required"`
	Description string `binding:"required"`
}

func Create(post *CreateParam) error {

	result := dborm.Db.Create(&dborm.Secret{
		UserId:      post.UserId,
		SecretId:    post.SecretId,
		SecretKey:   post.SecretKey,
		Description: post.Description,
	})

	return result.Error

}

// 更新密钥

type UpdateParam struct {
	Id          uint   `binding:"required"`
	UserId      uint   `binding:"required"`
	SecretId    string `binding:"required"`
	SecretKey   string `binding:"required"`
	Description string `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Secret{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Secret{
			SecretId:    post.SecretId,
			SecretKey:   post.SecretKey,
			Description: post.Description,
		})

	return result.Error

}

// 获取密钥列表

func FetchAll(userId uint) ([]*dborm.Secret, error) {

	var items []*dborm.Secret

	result := dborm.Db.Where(&dborm.Secret{UserId: userId}).Find(&items)

	return items, result.Error

}

// 获取密钥

func Fetch(id, userId uint) (*dborm.Secret, error) {

	var item *dborm.Secret

	result := dborm.Db.Where(&dborm.Secret{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除密钥

func Delete(id, userId uint) error {

	var item *dborm.Secret

	result := dborm.Db.Where(&dborm.Secret{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
