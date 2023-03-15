package passport

import (
	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/strutil"
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/model/user"
)

// 密钥迁移

func secretMigrator(u *dborm.User) (*dborm.User, error) {

	if u.AppKey != "" {
		// 新版密钥
		if _, err := strutil.Des3Decrypt(u.AppKey, args.Dataset.Secret); err == nil {
			return u, nil
		}
		// 旧版密钥
		if ak, err := strutil.Des3Decrypt(u.AppKey, u.Password); err == nil {
			user.Update(&user.UpdateParam{
				Id:       u.Id,
				AppKey:   ak,
				StoreKey: args.Dataset.Secret,
			})
		} else {
			return nil, err
		}
	} else {
		user.Update(&user.UpdateParam{
			Id:       u.Id,
			AppKey:   strutil.Rand(32),
			StoreKey: args.Dataset.Secret,
		})
	}

	return user.Fetch(&user.FetchParam{
		Id: u.Id,
	})

}
