package dborm

// 配置

type Config struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex:idx_config"`
	Value       string
	Module      string `gorm:"uniqueIndex:idx_config"`
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
	NSList      string
	Model       string
	CloudId     string `gorm:"uniqueIndex"`
	CloudMeta   any    `gorm:"serializer:json"`
	Description string
	Status      uint
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
	OSType      string
	Region      string
	Model       string
	CloudId     string `gorm:"uniqueIndex"`
	CloudMeta   any    `gorm:"serializer:json"`
	Description string
	Status      uint
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
	Request   any `gorm:"serializer:json"`
	Response  any `gorm:"serializer:json"`
	Status    string
	CreatedAt int64
	UpdatedAt int64
}

// 任务脚本

type TaskScript struct {
	Id            uint `gorm:"primaryKey"`
	UserId        uint `gorm:"index"`
	Name          string
	CommandType   string
	Username      string
	WorkDirectory string
	Content       string
	Description   string
	Timeout       uint
	CreatedAt     int64
	UpdatedAt     int64
}

// 用户

type User struct {
	Id          uint   `gorm:"primaryKey"`
	AppId       string `gorm:"uniqueIndex"`
	Username    string `gorm:"uniqueIndex"`
	Password    string
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
