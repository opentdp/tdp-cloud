package dborm

import (
	"time"

	"gorm.io/gorm"
)

// 公共模型

type TableModel struct {
	Id        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 用户表

type User struct {
	TableModel
	Username string `gorm:"index,unique"`
	Password string
	Secrets  []Secret
	Sessions []Session
}

// 密钥表

type Secret struct {
	TableModel
	UserId      int    `gorm:"index"`
	SecretId    string `gorm:"index,unique"`
	SecretKey   string
	Description string
}

// 会话表

type Session struct {
	TableModel
	UserId int    `gorm:"index"`
	Token  string `gorm:"index,unique"`
}
