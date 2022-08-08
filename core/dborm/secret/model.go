package secret

import (
	"tdp-cloud/core/dborm"
)

// 添加密钥

type CreateParam struct {
	UserId      uint   `json:"userId"`
	SecretId    string `json:"secretId" binding:"required"`
	SecretKey   string `json:"secretKey" binding:"required"`
	Description string `json:"description" binding:"required"`
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
	Id          uint   `json:"id"  binding:"required"`
	UserId      uint   `json:"userId" binding:"required"`
	SecretId    string `json:"secretId" binding:"required"`
	SecretKey   string `json:"secretKey" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.Model(&dborm.Secret{}).
		Where("id = ? AND user_id = ?", post.Id, post.UserId).
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

	result := dborm.Db.Find(&items, "user_id = ?", userId)

	return items, result.Error

}

// 获取密钥

func Fetch(id, userId uint) (dborm.Secret, error) {

	var item dborm.Secret

	result := dborm.Db.First(&item, "id = ? AND user_id = ?", id, userId)

	return item, result.Error

}

// 删除密钥

func Delete(id, userId uint) error {

	var item dborm.Secret

	result := dborm.Db.Delete(&item, "id = ? AND user_id = ?", id, userId)

	return result.Error

}
