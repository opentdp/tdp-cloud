package migrator

import (
	"tdp-cloud/helper/logman"
	"tdp-cloud/module/model/migration"
)

func Deploy() {

	if err := doMigrate(); err != nil {
		logman.Fatal("Migrate database error:", err)
	}

}

func addMigration(k, v string) error {

	_, err := migration.Create(&migration.CreateParam{
		Version: k, Description: v,
	})

	return err

}

func isMigrated(k string) bool {

	rq := &migration.FetchParam{Version: k}

	if rs, err := migration.Fetch(rq); err == nil {
		return rs.Id > 0
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

	if err := v100002(); err != nil {
		return err
	}

	return nil

}
