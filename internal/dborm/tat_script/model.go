package tat_script

import (
	"tdp-cloud/internal/dborm"
)

type CreateParam struct {
	UserId           uint
	Name             string `binding:"required"`
	Username         string `binding:"required"`
	Description      string
	Content          string `binding:"required"`
	CommandType      string `binding:"required"`
	WorkingDirectory string `binding:"required"`
	Timeout          uint   `binding:"required"`
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.TATScript{
		UserId:           post.UserId,
		Name:             post.Name,
		Username:         post.Username,
		Description:      post.Description,
		Content:          post.Content,
		WorkingDirectory: post.WorkingDirectory,
		CommandType:      post.CommandType,
		Timeout:          post.Timeout,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

type UpdateParam struct {
	Id               uint `binding:"required"`
	UserId           uint
	Name             string `binding:"required"`
	Description      string
	Content          string `binding:"required"`
	Username         string `binding:"required"`
	CommandType      string `binding:"required"`
	WorkingDirectory string `binding:"required"`
	Timeout          uint   `binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.TATScript{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.TATScript{
			Name:             post.Name,
			Description:      post.Description,
			Content:          post.Content,
			Username:         post.Username,
			CommandType:      post.CommandType,
			WorkingDirectory: post.WorkingDirectory,
			Timeout:          post.Timeout,
		})

	return result.Error

}

func FetchAll(userId uint) ([]*dborm.TATScript, error) {

	var items []*dborm.TATScript

	result := dborm.Db.Where(&dborm.TATScript{UserId: userId}).Find(&items)

	return items, result.Error

}

func Fetch(id, userId uint) (*dborm.TATScript, error) {

	var item *dborm.TATScript

	result := dborm.Db.Where(&dborm.TATScript{Id: id, UserId: userId}).Find(&item)

	return item, result.Error

}

func Delete(id, userId uint) error {

	var item *dborm.TATScript

	result := dborm.Db.Where(&dborm.TATScript{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
