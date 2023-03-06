package passport

import (
	"tdp-cloud/helper/strutil"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/model/user"
)

// 密钥迁移

func SecretMigrator(id uint, password string) (*dborm.User, error) {

	err := user.Update(&user.UpdateParam{
		Id:       id,
		Password: password,
		AppKey:   strutil.Rand(32),
	})

	if err != nil {
		return nil, err
	}

	return user.Fetch(&user.FetchParam{Id: id})

}
