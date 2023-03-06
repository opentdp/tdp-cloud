package migrator

import (
	"tdp-cloud/module/dborm"
)

func v100000() error {

	return dborm.Db.AutoMigrate(
		&dborm.Certjob{},
		&dborm.Config{},
		&dborm.Cronjob{},
		&dborm.Domain{},
		&dborm.Keypair{},
		&dborm.Machine{},
		&dborm.Script{},
		&dborm.Taskline{},
		&dborm.User{},
		&dborm.Vendor{},
	)

}
