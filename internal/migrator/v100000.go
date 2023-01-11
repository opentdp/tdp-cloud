package migrator

import (
	"tdp-cloud/internal/dborm"
)

func v100001() error {

	return dborm.Db.AutoMigrate(
		&dborm.Config{},
		&dborm.Domain{},
		&dborm.Machine{},
		&dborm.Session{},
		&dborm.Sshkey{},
		&dborm.TaskHistory{},
		&dborm.TaskScript{},
		&dborm.User{},
		&dborm.Vendor{},
	)

}
