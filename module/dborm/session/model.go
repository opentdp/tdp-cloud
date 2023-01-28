package session

import (
	"errors"
	"time"

	"tdp-cloud/helper/strings"
	"tdp-cloud/module/dborm"
)

// 添加令牌

func Create(userId uint) (string, error) {

	item := &dborm.Session{
		UserId: userId,
		Token:  strings.Rand(32),
	}

	result := dborm.Db.Create(item)

	return item.Token, result.Error

}

// 获取令牌

func Fetch(token string) (*dborm.Session, error) {

	var item *dborm.Session

	dborm.Db.Where(&dborm.Session{Token: token}).First(&item)

	if item.Id == 0 {
		return nil, errors.New("会话不存在")
	}

	// 会话超过30分钟，删除令牌
	if time.Now().Unix()-item.UpdatedAt > 1800 {
		dborm.Db.Delete(&item)
		return nil, errors.New("会话已过期")
	}

	// 会话超过1分钟，自动续期
	if time.Now().Unix()-item.UpdatedAt > 60 {
		dborm.Db.Save(&item)
	}

	return item, nil

}
