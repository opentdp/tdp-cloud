package session

import (
	"errors"
	"time"

	"tdp-cloud/helper/strutil"
	"tdp-cloud/module/dborm"
)

// 创建会话

type CreateParam struct {
	UserId    uint
	UserAgent string
}

func Create(data *CreateParam) (string, error) {

	item := &dborm.Session{
		UserId:    data.UserId,
		UserAgent: data.UserAgent,
		Token:     strutil.Rand(32),
	}

	result := dborm.Db.Create(item)

	return item.Token, result.Error

}

// 更新会话

type UpdateParam struct {
	Id        uint `binding:"required"`
	UserId    uint
	UserAgent string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&dborm.Session{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Updates(dborm.Session{
			UserAgent: data.UserAgent,
		})

	return result.Error

}

// 删除会话

type DeleteParam struct {
	Id     uint
	UserId uint
}

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&dborm.Session{
			Id:     data.Id,
			UserId: data.UserId,
		}).
		Delete(&dborm.Session{})

	return result.Error

}

// 获取会话

type FetchParam struct {
	Id    uint
	Token string
}

func Fetch(data *FetchParam) (*dborm.Session, error) {

	var item *dborm.Session

	dborm.Db.
		Where(&dborm.Session{
			Id:    data.Id,
			Token: data.Token,
		}).
		First(&item)

	if item.Id == 0 {
		return nil, errors.New("会话不存在")
	}

	// 会话停留时长
	stay := time.Now().Unix() - item.UpdatedAt

	// 停留超过60分钟，删除会话
	if stay > 3600 {
		dborm.Db.Delete(&item)
		return nil, errors.New("会话已过期")
	}

	// 停留超过5分钟，自动续期
	if stay > 300 {
		dborm.Db.Save(&item)
	}

	return item, nil

}

// 获取会话列表

type FetchAllParam struct {
	UserId uint
}

func FetchAll(data *FetchAllParam) ([]*dborm.Session, error) {

	var items []*dborm.Session

	result := dborm.Db.
		Where(&dborm.Session{
			UserId: data.UserId,
		}).
		Find(&items)

	return items, result.Error

}

// 获取会话总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Where(&dborm.Session{
			UserId: data.UserId,
		}).
		Count(&count)

	return count, result.Error

}
