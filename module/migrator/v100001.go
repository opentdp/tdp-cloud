package migrator

import (
	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/strutil"
	"tdp-cloud/module/model/config"
	"tdp-cloud/module/model/user"
)

func v100001() error {

	// TODO: 临时处理
	v100000MigrateLog()

	if isMigrated("v100001") {
		return nil
	}

	if err := v100001AddUser(); err != nil {
		return err
	}

	return addMigration("v100001", "添加默认账号")

}

func v100001AddUser() error {

	_, err := user.Create(&user.CreateParam{
		Username: "admin",
		Password: "123456",
		Level:    1,
		AppKey:   strutil.Rand(32),
		Email:    "admin@tdp.icu",
		StoreKey: args.Dataset.Secret,
	})

	return err

}

func v100000MigrateLog() {

	res, err := config.Fetch(&config.FetchParam{Name: "v100001"})

	if err == nil && res.Id > 0 {
		config.Delete(&config.DeleteParam{Id: res.Id})
		addMigration("v100001", "添加默认账号")
	}

}
