package dborm

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(dsn string) {

	var err error

	Db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	if Db.AutoMigrate(&User{}, &Secret{}, &Session{}) != nil {
		panic("Failed to migrate database")
	}

}
