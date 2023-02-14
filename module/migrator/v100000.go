package migrator

import (
	"tdp-cloud/module/dborm"
)

func v100000() error {

	return dborm.Db.AutoMigrate(
		&dborm.Config{},
		&dborm.Cronjob{},
		&dborm.Domain{},
		&dborm.Keypair{},
		&dborm.Machine{},
		&dborm.Session{},
		&dborm.Taskline{},
		&dborm.Script{},
		&dborm.User{},
		&dborm.Vendor{},
	)

}
