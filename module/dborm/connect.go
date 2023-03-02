package dborm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"tdp-cloud/helper/logman"
)

var Db *gorm.DB

func Connect() {

	config := &gorm.Config{
		Logger: NewLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if db, err := gorm.Open(dialector(), config); err != nil {
		logman.Fatal("Connect to databse error:", err)
	} else {
		Db = db
	}

}
