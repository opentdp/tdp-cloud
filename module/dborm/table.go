package dborm

// 域名证书

type Certjob struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint `gorm:"index"`
	VendorId  uint `gorm:"index"`
	Email     string
	Domain    string `gorm:"uniqueIndex"`
	CaType    string
	EabKeyId  string
	EabMacKey string
	History   any `gorm:"serializer:json"`
	CreatedAt int64
	UpdatedAt int64
}

// 系统配置

type Config struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex:idx_config"`
	Value       string
	Module      string `gorm:"uniqueIndex:idx_config"`
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 计划任务

type Cronjob struct {
	Id         uint `gorm:"primaryKey"`
	UserId     uint `gorm:"index"`
	Name       string
	Type       string
	Content    string
	Second     string
	Minute     string
	Hour       string
	DayofMonth string
	Month      string
	DayofWeek  string
	Location   string
	PrevTime   int64
	NextTime   int64
	CreatedAt  int64
	UpdatedAt  int64
}

// 域名资源

type Domain struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	VendorId    uint `gorm:"index"`
	Name        string
	NSList      string
	Model       string
	CloudId     string `gorm:"uniqueIndex"`
	CloudMeta   any    `gorm:"serializer:json"`
	Status      string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 密钥对

type Keypair struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	PublicKey   string
	PrivateKey  string `json:"-"`
	KeyType     string `gorm:"index"`
	Cipher      string
	Status      string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 主机资源

type Machine struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	VendorId    uint `gorm:"index"`
	HostName    string
	IpAddress   string
	OSType      string
	Region      string
	Model       string
	CloudId     string `gorm:"uniqueIndex,default:null"`
	CloudMeta   any    `gorm:"serializer:json"`
	WorkerId    string `gorm:"uniqueIndex,default:null"`
	WorkerMeta  any    `gorm:"serializer:json"`
	Status      string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
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
	Timeout       uint
	Description   string
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
	Id          uint     `gorm:"primaryKey"`
	Username    string   `gorm:"uniqueIndex"`
	Password    string   `json:"-"`
	Level       uint     `gorm:"default:5"`
	AppId       string   `gorm:"uniqueIndex"`
	AppKey      string   `json:"-"`
	Email       string   `gorm:"uniqueIndex,default:null"`
	Description string   `gorm:"default:挥一挥手"`
	Vendors     []Vendor `json:",omitempty"`
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
	Cipher      string
	Status      string
	Description string
	Certjobs    []Certjob `json:",omitempty"`
	Domains     []Domain  `json:",omitempty"`
	Machines    []Machine `json:",omitempty"`
	CreatedAt   int64
	UpdatedAt   int64
}
