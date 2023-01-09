package dborm

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
	Status      string
	CloudData   string
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
	Address     string
	Status      string
	CloudData   string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 自动化助手 - 脚本

type TATScript struct {
	Id               uint `gorm:"primaryKey"`
	UserId           uint `gorm:"index"`
	Name             string
	Username         string
	Content          string
	Description      string
	CommandType      string
	WorkingDirectory string
	Timeout          uint
	CreatedAt        int64
	UpdatedAt        int64
}

// 自动化助手 - 历史记录

type TATHistory struct {
	Id                   uint `gorm:"primaryKey"`
	UserId               uint `gorm:"index"`
	VendorId             uint `gorm:"index"`
	Name                 string
	Region               string
	InvocationId         string
	InvocationStatus     string
	InvocationResultJson string
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

// 子节点任务

type Worktask struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index"`
	HostId    string `gorm:"index"`
	HostName  string
	Subject   string
	Status    string
	Request   string
	Response  string
	CreatedAt int64
	UpdatedAt int64
}
