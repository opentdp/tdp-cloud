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

// 密钥对

type Keypair struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	PublicKey   string
	PrivateKey  string
	KeyType     uint `gorm:"index"`
	Description string
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
	WorkerId    string `gorm:"uniqueIndex"`
	WorkerMeta  any    `gorm:"serializer:json"`
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

// 命令脚本

type Script struct {
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

// 任务记录

type Taskline struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint `gorm:"index"`
	Subject   string
	HostName  string
	WorkerId  string `gorm:"index"`
	Request   any    `gorm:"serializer:json"`
	Response  any    `gorm:"serializer:json"`
	Status    string
	CreatedAt int64
	UpdatedAt int64
}

// 用户

type User struct {
	Id          uint   `gorm:"primaryKey"`
	AppId       string `gorm:"uniqueIndex"`
	Username    string `gorm:"uniqueIndex"`
	Password    string `json:"-"`
	Level       uint   `gorm:"default:5"`
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
	SecretKey   string `json:"-"`
	Provider    string
	Description string
	Domains     []Domain
	Machines    []Machine
	CreatedAt   int64
	UpdatedAt   int64
}
