package dborm

import (
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(dsn string) {

	var err error

	if strings.Index(dsn, "@") > 0 {
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		Db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect database")
	}

	if Db.AutoMigrate(&User{}, &Secret{}, &Session{}, &TAT{}) != nil {
		panic("Failed to migrate database")
	}

}
