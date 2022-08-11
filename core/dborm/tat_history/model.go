package tat_history

import (
	"tdp-cloud/core/dborm"
)

type CreateParam struct {
	Name         string `binding:"required"`
	InvocationId string `binding:"required"`
	Region       string `binding:"required"`
	UserId       uint
	KeyId        uint
}

func Create(post *CreateParam) error {

	result := dborm.Db.Create(&dborm.TATHistory{
		UserId:               post.UserId,
		KeyId:                post.KeyId,
		Name:                 post.Name,
		Region:               post.Region,
		InvocationId:         post.InvocationId,
		InvocationStatus:     "",
		InvocationResultJson: "",
	})

	return result.Error

}

type UpdateParam struct {
	Id                   uint
	InvocationStatus     string `binding:"required"`
	InvocationResultJson string `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.Model(&dborm.TATHistory{}).
		Where("id = ?", post.Id).
		Updates(dborm.TATHistory{
			InvocationStatus:     post.InvocationStatus,
			InvocationResultJson: post.InvocationResultJson,
		})

	return result.Error

}

func FetchAll(userId, keyId uint) ([]*dborm.TATHistory, error) {

	var items []*dborm.TATHistory

	result := dborm.Db.Limit(50).Order("id desc").
		Find(&items, "user_id = ? and key_id = ?", userId, keyId)

	return items, result.Error

}

func Delete(id int) error {

	result := dborm.Db.Delete(&dborm.TATHistory{}, id)

	return result.Error

}
