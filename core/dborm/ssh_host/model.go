package item

import (
	"tdp-cloud/core/dborm"
)

// 添加主机

type CreateParam struct {
	UserId      uint   `json:"userId"`
	Address     string `json:"Address" binding:"required"`
	Username    string `json:"Username" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func Create(post *CreateParam) error {

	result := dborm.Db.Create(&dborm.SSHHost{
		UserId:      post.UserId,
		Address:     post.Address,
		Username:    post.Username,
		Description: post.Description,
	})

	return result.Error

}

// 更新主机

type UpdateParam struct {
	Id          uint   `json:"id"  binding:"required"`
	UserId      uint   `json:"userId" binding:"required"`
	Address     string `json:"Address" binding:"required"`
	Username    string `json:"Username" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.Model(&dborm.SSHHost{}).
		Where("id = ? AND user_id = ?", post.Id, post.UserId).
		Updates(dborm.SSHHost{
			Address:     post.Address,
			Username:    post.Username,
			Description: post.Description,
		})

	return result.Error

}

// 获取主机列表

func FetchAll(userId uint) ([]*dborm.SSHHost, error) {

	var items []*dborm.SSHHost

	result := dborm.Db.Find(&items, "user_id = ?", userId)

	return items, result.Error

}

// 获取主机

func Fetch(id, userId uint) (dborm.SSHHost, error) {

	var item dborm.SSHHost

	result := dborm.Db.First(&item, "id = ? AND user_id = ?", id, userId)

	return item, result.Error

}

// 删除主机

func Delete(id, userId uint) error {

	var item dborm.SSHHost

	result := dborm.Db.Delete(&item, "id = ? AND user_id = ?", id, userId)

	return result.Error

}
