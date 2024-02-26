package certjob

import (
	"github.com/opentdp/go-helper/dborm"

	"tdp-cloud/model"
)

// 创建计划

type CreateParam struct {
	UserId    uint
	VendorId  uint   `binding:"required"`
	Email     string `binding:"required"`
	Domain    string `binding:"required"`
	CaType    string `binding:"required"`
	EabKeyId  string
	EabMacKey string
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Certjob{
		UserId:    data.UserId,
		VendorId:  data.VendorId,
		Email:     data.Email,
		Domain:    data.Domain,
		CaType:    data.CaType,
		EabKeyId:  data.EabKeyId,
		EabMacKey: data.EabMacKey,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新计划

type UpdateParam struct {
	Id        uint
	UserId    uint
	VendorId  uint
	Email     string
	Domain    string
	CaType    string
	EabKeyId  string
	EabMacKey string
	History   any
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Certjob{
			Id:     data.Id,
			UserId: data.UserId,
			Domain: data.Domain,
		}).
		Updates(model.Certjob{
			VendorId:  data.VendorId,
			Email:     data.Email,
			Domain:    data.Domain,
			CaType:    data.CaType,
			EabKeyId:  data.EabKeyId,
			EabMacKey: data.EabMacKey,
			History:   data.History,
		})

	return result.Error

}

// 删除计划

type DeleteParam struct {
	Id     uint
	UserId uint
	Domain string
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&model.Certjob{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&model.Certjob{})

	return result.Error

}

// 获取计划

type FetchParam struct {
	Id     uint
	UserId uint
	Domain string
}

func Fetch(data *FetchParam) (*model.Certjob, error) {

	var item *model.Certjob

	result := dborm.Db.
		Where(&model.Certjob{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	return item, result.Error

}

// 获取计划列表

type FetchAllParam struct {
	UserId   uint
	VendorId uint
}

func FetchAll(data *FetchAllParam) ([]*model.Certjob, error) {

	var items []*model.Certjob

	result := dborm.Db.
		Where(&model.Certjob{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取计划总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.Certjob{}).
		Where(&model.Certjob{
			UserId:   data.UserId,
			VendorId: data.VendorId,
		}).
		Count(&count)

	return count, result.Error

}
