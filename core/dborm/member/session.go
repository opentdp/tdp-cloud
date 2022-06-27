package member

import (
	"tdp-cloud/core/dborm"
)

// 获取会话

func FetchSession(token string) dborm.Session {

	var session dborm.Session

	dborm.Db.First(&session, "token = ?", token)

	return session

}
