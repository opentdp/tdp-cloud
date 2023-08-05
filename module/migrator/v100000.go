package migrator

import (
	"github.com/opentdp/go-helper/dborm"

	"tdp-cloud/model"
)

func v100000() error {

	return v100000AutoMigrate()

}

func v100000AutoMigrate() error {

	// NOTE: 有外键的表需要先导入
	return dborm.Db.AutoMigrate(
		&model.User{},
		&model.Vendor{},
		&model.Certjob{},
		&model.Config{},
		&model.Cronjob{},
		&model.Domain{},
		&model.Keypair{},
		&model.Machine{},
		&model.Migration{},
		&model.Script{},
		&model.Taskline{},
	)

}
