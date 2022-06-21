package user

import (
	"tdp-cloud/core/dborm"
)

// 添加密钥

type SecretParam struct {
	UserId      uint   `json:"userId"`
	SecretId    string `json:"secretId" binding:"required"`
	SecretKey   string `json:"secretKey" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateSecret(post *SecretParam) error {

	result := dborm.Db.Create(&dborm.Secret{
		UserId:      post.UserId,
		SecretId:    post.SecretId,
		SecretKey:   post.SecretKey,
		Description: post.Description,
	})

	return result.Error

}

// 删除密钥

func DeleteSecret(userId, id uint) error {

	var secret dborm.Secret

	result := dborm.Db.Delete(&secret, "user_id = ? AND id = ?", userId, id)

	return result.Error

}

// 获取密钥列表

func FindSecrets(userId uint) []*dborm.Secret {

	var secrets []*dborm.Secret

	dborm.Db.Find(&secrets, "user_id = ?", userId)

	return secrets

}

// 获取密钥

func FetchSecret(userId, id uint) dborm.Secret {

	var secret dborm.Secret

	dborm.Db.First(&secret, "user_id = ? AND id = ?", userId, id)

	return secret

}
