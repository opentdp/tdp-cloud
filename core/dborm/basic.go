package dborm

import (
	"os"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func Connect(dsn string) {

	var err error

	var config = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if os.Getenv("IS_DEBUG") != "" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	if strings.Index(dsn, "@") > 0 {
		Db, err = gorm.Open(mysql.Open(dsn), config)
	} else {
		Db, err = gorm.Open(sqlite.Open(dsn), config)
	}

	if err != nil {
		panic("Failed to connect database")
	}

	if Db.AutoMigrate(&User{}, &Secret{}, &Session{}, &TAT{}) != nil {
		panic("Failed to migrate database")
	}

}
