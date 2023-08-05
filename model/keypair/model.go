package keypair

import (
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/go-helper/secure"

	"tdp-cloud/model"
)

// 创建密钥

type CreateParam struct {
	UserId      uint
	PublicKey   string `binding:"required"`
	PrivateKey  string `binding:"required"`
	KeyType     string `binding:"required"`
	Cipher      string
	Status      string
	Description string `binding:"required"`
	StoreKey    string // 存储密钥
}

func Create(data *CreateParam) (uint, error) {

	if data.PrivateKey != "" && data.StoreKey != "" {
		secret, err := secure.Des3Encrypt(data.PrivateKey, data.StoreKey)
		if err == nil {
			data.PrivateKey = secret
			data.Cipher = "appkey"
		}
	}

	item := &model.Keypair{
		UserId:      data.UserId,
		PublicKey:   data.PublicKey,
		PrivateKey:  data.PrivateKey,
		KeyType:     data.KeyType,
		Cipher:      data.Cipher,
		Status:      data.Status,
		Description: data.Description,
	}

	result := dborm.Db.Create(item)

	return item.Id, result.Error

}

// 更新密钥

type UpdateParam struct {
	Id          uint
	UserId      uint
	PublicKey   string
	PrivateKey  string
	KeyType     string
	Cipher      string
	Status      string
	Description string
	StoreKey    string // 存储密钥
}

func Update(data *UpdateParam) error {

	if data.PrivateKey != "" && data.StoreKey != "" {
		secret, err := secure.Des3Encrypt(data.PrivateKey, data.StoreKey)
		if err == nil {
			data.PrivateKey = secret
			data.Cipher = "appkey"
		}
	}

	result := dborm.Db.
		Where(&model.Keypair{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(model.Keypair{
			PublicKey:   data.PublicKey,
			PrivateKey:  data.PrivateKey,
			KeyType:     data.KeyType,
			Cipher:      data.Cipher,
			Status:      data.Status,
			Description: data.Description,
		})

	return result.Error

}

// 删除密钥

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&model.Keypair{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&model.Keypair{})

	return result.Error

}

// 获取密钥

type FetchParam struct {
	Id       uint
	UserId   uint
	StoreKey string // 存储密钥
}

func Fetch(data *FetchParam) (*model.Keypair, error) {

	var item *model.Keypair

	result := dborm.Db.
		Where(&model.Keypair{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		First(&item)

	if item.Cipher != "" && data.StoreKey != "" {
		item.PrivateKey, _ = secure.Des3Decrypt(item.PrivateKey, data.StoreKey)
	}

	return item, result.Error

}

// 获取密钥列表

type FetchAllParam struct {
	UserId  uint
	KeyType string
}

func FetchAll(data *FetchAllParam) ([]*model.Keypair, error) {

	var items []*model.Keypair

	result := dborm.Db.
		Where(&model.Keypair{
			UserId:  data.UserId,
			KeyType: data.KeyType,
		}).
		Find(&items)

	return items, result.Error

}

// 获取密钥总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&model.Keypair{}).
		Where(&model.Keypair{
			UserId:  data.UserId,
			KeyType: data.KeyType,
		}).
		Count(&count)

	return count, result.Error

}
