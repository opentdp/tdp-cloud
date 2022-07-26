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

// 删除密钥

func Delete(userId, id uint) error {

	var secret dborm.Secret

	result := dborm.Db.Delete(&secret, "user_id = ? AND id = ?", userId, id)

	return result.Error

}

// 获取密钥列表

func Find(userId uint) ([]*dborm.Secret, error) {

	var secrets []*dborm.Secret

	result := dborm.Db.Find(&secrets, "user_id = ?", userId)

	return secrets, result.Error

}

// 获取密钥

func Fetch(userId, id uint) (dborm.Secret, error) {

	var secret dborm.Secret

	result := dborm.Db.First(&secret, "user_id = ? AND id = ?", userId, id)

	return secret, result.Error

}
