package tat_history

import (
	"tdp-cloud/internal/dborm"
)

// 添加历史

type CreateParam struct {
	UserId       uint
	KeyId        uint
	Name         string `binding:"required"`
	Region       string `binding:"required"`
	InvocationId string `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.TATHistory{
		UserId:               post.UserId,
		KeyId:                post.KeyId,
		Name:                 post.Name,
		Region:               post.Region,
		InvocationId:         post.InvocationId,
		InvocationStatus:     "",
		InvocationResultJson: "",
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新历史

type UpdateParam struct {
	Id                   uint
	UserId               uint
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

// 获取历史列表

func FetchAll(userId, keyId uint) ([]*dborm.TATHistory, error) {

	var items []*dborm.TATHistory

	result := dborm.Db.
		Where(&dborm.TATHistory{UserId: userId, KeyId: keyId}).
		Limit(50).Order("id DESC").
		Find(&items)

	return items, result.Error

}

// 获取历史

func Fetch(id, userId uint) (*dborm.TATHistory, error) {

	var item *dborm.TATHistory

	result := dborm.Db.Where(&dborm.TATHistory{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除历史

func Delete(id, userId uint) error {

	var item *dborm.Secret

	result := dborm.Db.Where(&dborm.TATHistory{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
