package tat_history

import (
	"tdp-cloud/core/dborm"
)

type CreateParam struct {
	UserId       uint
	KeyId        uint
	Name         string `binding:"required"`
	Region       string `binding:"required"`
	InvocationId string `binding:"required"`
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
	InvocationId         string
	InvocationStatus     string `binding:"required"`
	InvocationResultJson string `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.TATHistory{Id: post.Id, InvocationId: post.InvocationId}).
		Updates(&dborm.TATHistory{
			InvocationStatus:     post.InvocationStatus,
			InvocationResultJson: post.InvocationResultJson,
		})

	return result.Error

}

func FetchAll(userId, keyId uint) ([]*dborm.TATHistory, error) {

	var items []*dborm.TATHistory

	result := dborm.Db.
		Where(&dborm.TATHistory{UserId: userId, KeyId: keyId}).
		Limit(50).Order("id DESC").
		Find(&items)

	return items, result.Error

}

type FetchParam struct {
	Id           uint
	InvocationId string
}

func Fetch(post *FetchParam) (*dborm.TATHistory, error) {

	var item *dborm.TATHistory

	result := dborm.Db.
		Where(&dborm.TATHistory{Id: post.Id, InvocationId: post.InvocationId}).
		First(&item)

	return item, result.Error

}

func Delete(id int) error {

	result := dborm.Db.Delete(&dborm.TATHistory{}, id)

	return result.Error

}
