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
	IpAddress   string `binding:"required"`
	Region      string
	Model       string `binding:"required"`
	CloudId     string
	CloudMeta   string
	Description string
	Status      string
}

func Create(post *CreateParam) (uint, error) {

	item := &dborm.Machine{
		UserId:      post.UserId,
		VendorId:    post.VendorId,
		HostName:    post.HostName,
		IpAddress:   post.IpAddress,
		Region:      post.Region,
		Model:       post.Model,
		CloudId:     post.CloudId,
		CloudMeta:   datatypes.JSON(post.CloudMeta),
		Description: post.Description,
		Status:      datatypes.JSON(post.Status),
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
	IpAddress   string `binding:"required"`
	Region      string
	Model       string `binding:"required"`
	CloudId     string
	CloudMeta   string
	Description string
	Status      string
}

func Update(post *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Machine{Id: post.Id, UserId: post.UserId}).
		Updates(dborm.Machine{
			VendorId:    post.VendorId,
			HostName:    post.HostName,
			IpAddress:   post.IpAddress,
			Region:      post.Region,
			Model:       post.Model,
			CloudId:     post.CloudId,
			CloudMeta:   datatypes.JSON(post.CloudMeta),
			Description: post.Description,
			Status:      datatypes.JSON(post.Status),
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
