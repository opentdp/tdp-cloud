package migrator

import (
	"tdp-cloud/internal/dborm/user"
)

func v100001() error {

	if isMigrated("v100001") {
		return nil
	}

	if err := initAdminUser(); err != nil {
		return err
	}

	return addMigration("v100001", "添加默认账号")

}

func initAdminUser() error {

	item, err := user.Fetch(&user.FetchParam{Id: 1})

	if err == nil && item.Id > 0 {
		return nil
	}

	_, err = user.Create(&user.CreateParam{
		Username: "admin",
		Password: "123456",
	})

	return err

}
