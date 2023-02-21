package dborm

import (
	"log"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func Connect(dsn string) {

	var err error
	var logLevel = logger.Silent

	if viper.GetBool("debug") {
		logLevel = logger.Info
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
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
