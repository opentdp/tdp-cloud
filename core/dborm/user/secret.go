package user

import (
	"tdp-cloud/core/dborm"
)

// 添加密钥

func CreateSecret(post *SecretInput) error {

	result := dborm.Db.Create(&dborm.Secret{
		UserId:      post.UserId,
		SecretId:    post.SecretId,
		SecretKey:   post.SecretKey,
		Description: post.Description,
	})

	return result.Error

}

// 删除密钥

func DeleteSecret(userId, id int) error {

	var item dborm.Secret

	result := dborm.Db.Delete(&item, "user_id = ? AND id = ?", userId, id)

	return result.Error

}

// 获取密钥列表

func FindSecrets(userId int) []*dborm.Secret {

	var list []*dborm.Secret

	dborm.Db.Find(&list, "user_id = ?", userId)

	return list

}

// 获取密钥

func FetchSecret(userId, id int) dborm.Secret {

	var item dborm.Secret

	dborm.Db.First(&item, "user_id = ? AND id = ?", userId, id)

	return item

}
