package migrator

import (
	"tdp-cloud/module/dborm"
)

func v100000() error {

	if err := v100000AutoMigrate(); err != nil {
		return err
	}

	return nil

}

func v100000AutoMigrate() error {

	return dborm.Db.AutoMigrate(
		&dborm.Certjob{},
		&dborm.Config{},
		&dborm.Cronjob{},
		&dborm.Domain{},
		&dborm.Keypair{},
		&dborm.Machine{},
		&dborm.Migration{},
		&dborm.Script{},
		&dborm.Taskline{},
		&dborm.User{},
		&dborm.Vendor{},
	)

}
