package dborm

import (
	"time"

	"gorm.io/gorm"
)

type TableModel struct {
	Id        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	TableModel
	Username string `gorm:"index,unique"`
	Password string
	Secrets  []Secret
	Sessions []Session
}

type Secret struct {
	TableModel
	UserId      int    `gorm:"index"`
	SecretId    string `gorm:"index,unique"`
	SecretKey   string
	Description string
}

type Session struct {
	TableModel
	UserId int    `gorm:"index"`
	Token  string `gorm:"index,unique"`
}
