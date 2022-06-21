package dborm

import (
	"time"

	"gorm.io/gorm"
)

// 公共模型

type TableModel struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 用户表

type User struct {
	TableModel
	Username    string `gorm:"index,unique"`
	Password    string
	Description string `gorm:"default:不可能！我的代码怎么可能会有bug！"`
	Secrets     []Secret
	Sessions    []Session
}

// 密钥表

type Secret struct {
	TableModel
	UserId      uint   `gorm:"index"`
	SecretId    string `gorm:"index,unique"`
	SecretKey   string
	Description string
}

// 会话表

type Session struct {
	TableModel
	UserId uint   `gorm:"index"`
	Token  string `gorm:"index,unique"`
}
