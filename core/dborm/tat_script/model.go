package tat_script

import (
	"tdp-cloud/core/dborm"
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

func Create(post *CreateParam) error {
	result := dborm.Db.Create(&dborm.TATScript{
		UserId:           post.UserId,
		Name:             post.Name,
		Username:         post.Username,
		Description:      post.Description,
		Content:          post.Content,
		WorkingDirectory: post.WorkingDirectory,
		CommandType:      post.CommandType,
		Timeout:          post.Timeout,
	})
	return result.Error
}

type UpdateParam struct {
	Id               uint   `binding:"required"`
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
		Model(&dborm.TATScript{}).
		Where("id = ?", post.Id).
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

func FetchAll(uid uint) ([]*dborm.TATScript, error) {
	var tatScripts []*dborm.TATScript

	result := dborm.Db.Find(&tatScripts, "user_id = ?", uid)
	return tatScripts, result.Error
}

func Fetch(id int) (dborm.TATScript, error) {
	var tatScript dborm.TATScript

	result := dborm.Db.Find(&tatScript, "id = ?", id)
	return tatScript, result.Error
}

func Delete(id int) error {
	result := dborm.Db.Delete(&dborm.TATScript{}, id)
	return result.Error
}
