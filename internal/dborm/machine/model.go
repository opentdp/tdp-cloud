package machine

import (
	"gorm.io/datatypes"

	"tdp-cloud/internal/dborm"
)

// 添加主机

type CreateParam struct {
	UserId      uint
	VendorId    uint   `binding:"required"`
	HostName    string `binding:"required"`
	Address     string `binding:"required"`
	Status      string
	CloudData   string `binding:"required"`
	Description string
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Machine{
		UserId:      post.UserId,
		VendorId:    post.VendorId,
		HostName:    post.HostName,
		Address:     post.Address,
		Status:      post.Status,
		CloudData:   datatypes.JSON(post.CloudData),
		Description: post.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新主机

type UpdateParam struct {
	Id          uint `binding:"required"`
	UserId      uint
	VendorId    uint   `binding:"required"`
	HostName    string `binding:"required"`
	Address     string `binding:"required"`
	Status      string
	CloudData   string `binding:"required"`
	Description string
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Machine{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Machine{
			VendorId:    post.VendorId,
			HostName:    post.HostName,
			Address:     post.Address,
			Status:      post.Status,
			CloudData:   datatypes.JSON(post.CloudData),
			Description: post.Description,
		})

	return result.Error

}

// 获取主机列表

func FetchAll(userId uint) ([]*dborm.Machine, error) {

	var items []*dborm.Machine

	result := dborm.Db.Where(&dborm.Machine{UserId: userId}).Find(&items)

	return items, result.Error

}

// 获取主机

func Fetch(id, userId uint) (*dborm.Machine, error) {

	var item *dborm.Machine

	result := dborm.Db.Where(&dborm.Machine{Id: id, UserId: userId}).First(&item)

	return item, result.Error

}

// 删除主机

func Delete(id, userId uint) error {

	var item *dborm.Machine

	result := dborm.Db.Where(&dborm.Machine{Id: id, UserId: userId}).Delete(&item)

	return result.Error

}
