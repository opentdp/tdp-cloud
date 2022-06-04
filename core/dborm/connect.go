package dborm

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "cloud.db"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	if db.AutoMigrate(&User{}, &Secret{}, &Session{}) != nil {
		panic("Failed to migrate database")
	}

	Db = db
}
