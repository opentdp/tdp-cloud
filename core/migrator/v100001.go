package migrator

import (
	"tdp-cloud/core/dborm"
)

func v00001() error {

	return dborm.Db.AutoMigrate(
		&dborm.Config{},
		&dborm.User{},
		&dborm.Session{},
		&dborm.Secret{},
		&dborm.Sshkey{},
		&dborm.TATScript{},
		&dborm.TATHistory{},
	)

}
