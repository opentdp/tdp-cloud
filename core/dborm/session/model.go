package session

import (
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

func FetchOne(token string) dborm.Session {

	var session dborm.Session

	dborm.Db.First(&session, "token = ?", token)

	// 会话超过30分钟，删除令牌
	if time.Now().Unix()-session.UpdatedAt > 1800 {
		dborm.Db.Delete(&session)
		return dborm.Session{}
	}

	// 会话超过1分钟，自动续期
	if time.Now().Unix()-session.UpdatedAt > 60 {
		dborm.Db.Save(&session)
	}

	return session

}
