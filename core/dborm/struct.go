package dborm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Secrets  []Secret
	Sessions []Session
}

type Secret struct {
	gorm.Model
	UserID    uint
	Describe  string
	SecretId  string
	SecretKey string
}

type Session struct {
	gorm.Model
	UserID uint
	Token  string
}
