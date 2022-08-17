package migrator

import (
	"tdp-cloud/internal/dborm/config"
	"tdp-cloud/internal/dborm/user"
)

func v100002() error {

	if isMigrated("100001") {
		return nil
	}

	if err := newAdminUser(); err != nil {
		return err
	}

	if err := newMigration(); err != nil {
		return err
	}

	return addMigration("100001")

}

func newAdminUser() error {

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

func newMigration() error {

	item, err := config.Fetch("Migration")

	if err == nil && item.Id > 0 {
		return nil
	}

	_, err = config.Create(&config.CreateParam{
		Name:        "Migration",
		Value:       Versions,
		Module:      "System",
		Description: "自动迁移记录",
	})

	return err

}
