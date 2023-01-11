package dborm

import (
	"gorm.io/datatypes"
)

// 配置

type Config struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex"`
	Value       string
	Module      string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 域名资源

type Domain struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	VendorId    uint
	Name        string
	Model       string
	CloudId     string `gorm:"uniqueIndex"`
	CloudMeta   datatypes.JSON
	Description string
	Status      datatypes.JSON
	CreatedAt   int64
	UpdatedAt   int64
}

// 主机资源

type Machine struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	VendorId    uint
	HostName    string
	IpAddress   string
	Region      string
	Model       string
	CloudId     string `gorm:"uniqueIndex"`
	CloudMeta   datatypes.JSON
	Description string
	Status      datatypes.JSON
	CreatedAt   int64
	UpdatedAt   int64
}

// 用户会话

type Session struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index"`
	Token     string `gorm:"uniqueIndex"`
	CreatedAt int64
	UpdatedAt int64
}

// SSH 密钥

type Sshkey struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	PublicKey   string
	PrivateKey  string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 任务记录

type TaskHistory struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index"`
	HostId    string `gorm:"index"`
	Subject   string
	HostName  string
	Request   datatypes.JSON
	Response  datatypes.JSON
	Status    string
	CreatedAt int64
	UpdatedAt int64
}

// 任务脚本

type TaskScript struct {
	Id            uint `gorm:"primaryKey"`
	UserId        uint `gorm:"index"`
	Name          string
	Username      string
	Description   string
	Content       string
	CommandType   string
	WorkDirectory string
	Timeout       uint
	CreatedAt     int64
	UpdatedAt     int64
}

// 用户

type User struct {
	Id          uint   `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex"`
	Password    string
	AppToken    string `gorm:"uniqueIndex"`
	Description string `gorm:"default:什么也没有"`
	Sessions    []Session
	Vendors     []Vendor
	CreatedAt   int64
	UpdatedAt   int64
}

// 厂商

type Vendor struct {
	Id          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"index"`
	SecretId    string `gorm:"uniqueIndex"`
	SecretKey   string
	Provider    string
	Description string
	Domains     []Domain
	Machines    []Machine
	CreatedAt   int64
	UpdatedAt   int64
}
