package dborm

import (
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"tdp-cloud/helper/logman"
)

var Db *gorm.DB

func Connect(dsn string) {

	var err error

	config := &gorm.Config{
		Logger: NewLogger(),
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
		logman.Fatal("Connect to databse error:", err)
	}

}
