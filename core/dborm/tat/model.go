package tat

import (
	"tdp-cloud/core/dborm"
)

type CreateParam struct {
	UserId      uint   `json:"userId"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Content     string `json:"content" binding:"required"`
}

func Create(post *CreateParam) error {

	result := dborm.Db.Create(&dborm.TAT{
		UserId:      post.UserId,
		Name:        post.Name,
		Description: post.Description,
		Content:     post.Content,
	})

	return result.Error

}

type UpdateParam struct {
	Id          uint   `json:"id"  binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

func Update(post *UpdateParam) error {
	result := dborm.Db.Model(&dborm.TAT{}).Where("id = ?", post.Id).Updates(dborm.TAT{Name: post.Name, Content: post.Content, Description: post.Description})
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
