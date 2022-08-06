package tat

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
	result := dborm.Db.Create(&dborm.TAT{
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
		Model(&dborm.TAT{}).
		Where("id = ?", post.Id).
		Updates(dborm.TAT{
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

func FetchAll(uid uint) ([]*dborm.TAT, error) {
	var tats []*dborm.TAT
	result := dborm.Db.Find(&tats, "user_id = ?", uid)
	return tats, result.Error
}

func FetchOne(id int) (dborm.TAT, error) {
	var tat dborm.TAT
	result := dborm.Db.Find(&tat, "id = ?", id)
	return tat, result.Error
}

func Delete(id int) error {
	result := dborm.Db.Delete(&dborm.TAT{}, id)
	return result.Error
}

type AddHistoryParam struct {
	Name         string `binding:"required"`
	InvocationId string `binding:"required"`
	Region       string `binding:"required"`

	UserId uint
	KeyId  uint
}

func AddHistory(post *AddHistoryParam) error {
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

type UpdateHistoryParam struct {
	Id                   uint
	InvocationStatus     string `binding:"required"`
	InvocationResultJson string `binding:"required"`
}

func UpdateHistory(post *UpdateHistoryParam) error {
	result := dborm.Db.Model(&dborm.TATHistory{}).
		Where("id = ?", post.Id).
		Updates(dborm.TATHistory{
			InvocationStatus:     post.InvocationStatus,
			InvocationResultJson: post.InvocationResultJson,
		})
	return result.Error
}

func FetchHistory(UserId, KeyId uint) ([]*dborm.TATHistory, error) {
	var tatHis []*dborm.TATHistory
	result := dborm.Db.Limit(50).Order("id desc").Find(&tatHis, "user_id = ? and key_id = ?", UserId, KeyId)
	return tatHis, result.Error
}

func DeleteHistory(id int) error {
	result := dborm.Db.Delete(&dborm.TATHistory{}, id)
	return result.Error
}
