package migrator

import (
	"log"

	"tdp-cloud/module/dborm/config"
)

func Deploy() {

	if err := doMigrate(); err != nil {
		log.Panic(err)
	}

}

func addMigration(k, v string) error {

	_, err := config.Create(&config.CreateParam{
		Name:        k,
		Value:       v,
		Module:      "Migration",
		Description: "数据库自动迁移记录",
	})

	return err

}

func isMigrated(k string) bool {

	if item, err := config.Fetch(k); err == nil {
		return item.Id > 0
	}

	return false

}

func doMigrate() error {

	if err := v100000(); err != nil {
		return err
	}

	if err := v100001(); err != nil {
		return err
	}

	return nil

}
