package session

import (
	"tdp-cloud/core/dborm"
	"tdp-cloud/core/utils"
	"time"
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

	// 会话已过期
	if session.UpdatedAt.Add(time.Minute * 15).Before(time.Now()) {
		dborm.Db.Delete(&session)
		return dborm.Session{}
	}

	dborm.Db.Save(&session)

	return session

}
