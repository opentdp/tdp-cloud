package keypair

import (
	"tdp-cloud/module/dborm"
)

// 添加密钥

type CreateParam struct {
	UserId      uint
	PublicKey   string `binding:"required"`
	PrivateKey  string `binding:"required"`
	KeyType     uint   `binding:"required"`
	Description string `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Keypair{
		UserId:      post.UserId,
		PublicKey:   post.PublicKey,
		PrivateKey:  post.PrivateKey,
		KeyType:     post.KeyType,
		Description: post.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新密钥

type UpdateParam struct {
	Id          uint `binding:"required"`
	UserId      uint
	PublicKey   string `binding:"required"`
	PrivateKey  string `binding:"required"`
	KeyType     uint
	Description string
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Keypair{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Keypair{
			PublicKey:   post.PublicKey,
			PrivateKey:  post.PrivateKey,
			KeyType:     post.KeyType,
			Description: post.Description,
		})

	return result.Error

}

// 获取密钥列表

func FetchAll(userId uint) ([]*dborm.Keypair, error) {

	var items []*dborm.Keypair

	result := dborm.Db.Where(&dborm.Keypair{UserId: userId}).Find(&items)

	return items, result.Error

}

// 获取密钥

func Fetch(id, userId uint) (*dborm.Keypair, error) {

	var item *dborm.Keypair

	result := dborm.Db.Where(&dborm.Keypair{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除密钥

func Delete(id, userId uint) error {

	var item *dborm.Keypair

	result := dborm.Db.Where(&dborm.Keypair{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
