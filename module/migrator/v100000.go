package migrator

import (
	"tdp-cloud/module/dborm"
	"tdp-cloud/module/model"
)

func v100000() error {

	if err := v100000AutoMigrate(); err != nil {
		return err
	}

	return nil

}

func v100000AutoMigrate() error {

	return dborm.Db.AutoMigrate(
		&model.Certjob{},
		&model.Config{},
		&model.Cronjob{},
		&model.Domain{},
		&model.Keypair{},
		&model.Machine{},
		&model.Migration{},
		&model.Script{},
		&model.Taskline{},
		&model.User{},
		&model.Vendor{},
	)

}
