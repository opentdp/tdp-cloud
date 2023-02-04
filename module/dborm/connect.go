package dborm

import (
	"log"
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
		Logger: logger.Default.LogMode(logger.Error),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if os.Getenv("TDP_DEBUG") != "" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	if strings.Contains(dsn, "@tcp") {
		if !strings.Contains(dsn, "?") {
			dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
		}
		Db, err = gorm.Open(mysql.Open(dsn), config)
	} else {
		if !strings.Contains(dsn, "?") {
			dsn += "?_pragma=busy_timeout=5000&_pragma=journa_mode(WAL)"
		}
		Db, err = gorm.Open(sqlite.Open(dsn), config)
	}

	if err != nil {
		log.Fatalln("Connect to databse error:", err)
	}

}
