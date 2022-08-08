package session

import (
	"errors"
	"time"

	"tdp-cloud/core/dborm"
	"tdp-cloud/core/utils"
)

// 添加令牌

func Create(userId uint) (string, error) {

	token := utils.RandString(32)

	result := dborm.Db.Create(&dborm.Session{
		UserId: userId,
		Token:  token,
	})

	return token, result.Error

}

// 获取令牌

func Fetch(token string) (*dborm.Session, error) {

	var item *dborm.Session

	result := dborm.Db.First(&item, "token = ?", token)

	if result.Error != nil {
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
