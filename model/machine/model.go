package machine

import (
	"github.com/opentdp/go-helper/dborm"

	"tdp-cloud/model"
)

// 创建主机

type CreateParam struct {
	UserId      uint
	VendorId    uint
	HostName    string `binding:"required"`
	IpAddress   string `binding:"required"`
	OSType      string `binding:"required"`
	Region      string
	Model       string `binding:"required"`
	CloudId     string
	CloudMeta   any
	WorkerId    string
	WorkerMeta  any
	Status      string
	Description string
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Machine{
		UserId:      data.UserId,
		VendorId:    data.VendorId,
		HostName:    data.HostName,
		IpAddress:   data.IpAddress,
		OSType:      data.OSType,
		Region:      data.Region,
		Model:       data.Model,
		CloudId:     data.CloudId,
		CloudMeta:   data.CloudMeta,
		WorkerId:    data.WorkerId,
		WorkerMeta:  data.WorkerMeta,
		Status:      data.Status,
		Description: data.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新主机

type UpdateParam struct {
	Id          uint
	UserId      uint
	VendorId    uint
	HostName    string
	IpAddress   string
	OSType      string
	Region      string
	Model       string
	CloudId     string
	CloudMeta   any
	WorkerId    string
	WorkerMeta  any
	Status      string
	Description string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Machine{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(model.Machine{
			VendorId:    data.VendorId,
			HostName:    data.HostName,
			IpAddress:   data.IpAddress,
			OSType:      data.OSType,
			Region:      data.Region,
			Model:       data.Model,
			CloudId:     data.CloudId,
			CloudMeta:   data.CloudMeta,
			WorkerId:    data.WorkerId,
			WorkerMeta:  data.WorkerMeta,
			Status:      data.Status,
			Description: data.Description,
		})

	return result.Error

}

// 删除主机

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&model.Machine{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&model.Machine{})

	return result.Error

}

// 获取主机

type FetchParam struct {
	Id       uint
	UserId   uint
	CloudId  string
	WorkerId string
}

func Fetch(data *FetchParam) (*model.Machine, error) {

	var item *model.Machine

	result := dborm.Db.
		Where(&model.Machine{
			Id:       data.Id,
			UserId:   data.UserId,
			CloudId:  data.CloudId,
			WorkerId: data.WorkerId,
		}).
		First(&item)

	return item, result.Error

}

// 获取主机列表

type FetchAllParam struct {
	UserId   uint
	VendorId uint
}

func FetchAll(data *FetchAllParam) ([]*model.Machine, error) {

	var items []*model.Machine

	result := dborm.Db.
		Where(&model.Machine{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取主机总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.Machine{}).
		Where(&model.Machine{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Count(&count)

	return count, result.Error

}
