package session

import (
	"errors"
	"time"

	"tdp-cloud/core/dborm"
	"tdp-cloud/core/helper"
)

// 添加令牌

func Create(userId uint) (string, error) {

	token := helper.RandString(32)

	result := dborm.Db.Create(&dborm.Session{
		UserId: userId,
		Token:  token,
	})

	return token, result.Error

}

// 获取令牌

func Fetch(token string) (*dborm.Session, error) {

	var item *dborm.Session

	result := dborm.Db.Where(&dborm.Session{Token: token}).First(&item)

	if result.Error != nil || item.Id == 0 {
		return nil, result.Error
	}

	// 会话超过30分钟，删除令牌
	if time.Now().Unix()-item.UpdatedAt > 1800 {
		dborm.Db.Delete(&item)
		return nil, errors.New("会话过期")
	}

	// 会话超过1分钟，自动续期
	if time.Now().Unix()-item.UpdatedAt > 60 {
		dborm.Db.Save(&item)
	}

	return item, nil

}
