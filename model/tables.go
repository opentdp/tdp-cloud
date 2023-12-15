package model

// 域名证书

type Certjob struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index"`
	VendorId  uint   `gorm:"index"`
	Email     string `gorm:"size:255"`
	Domain    string `gorm:"size:255;uniqueIndex"`
	CaType    string `gorm:"size:32"`
	EabKeyId  string `gorm:"size:128"`
	EabMacKey string `gorm:"size:128"`
	History   any    `gorm:"serializer:json"`
	CreatedAt int64
	UpdatedAt int64
}

// 系统配置

type Config struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:32;uniqueIndex:idx_config"`
	Value       string `gorm:"size:1024"`
	Type        string `gorm:"size:32;default:string"`
	Module      string `gorm:"size:32;uniqueIndex:idx_config"`
	Description string `gorm:"size:2048"`
	CreatedAt   int64
	UpdatedAt   int64
}

// 计划任务

type Cronjob struct {
	Id         uint   `gorm:"primaryKey"`
	UserId     uint   `gorm:"index"`
	Name       string `gorm:"size:128"`
	Type       string `gorm:"size:32"`
	Target     string `gorm:"size:32"`
	Content    string `gorm:"type:text"`
	Second     string `gorm:"size:32"`
	Minute     string `gorm:"size:32"`
	Hour       string `gorm:"size:32"`
	DayofMonth string `gorm:"size:32"`
	Month      string `gorm:"size:32"`
	DayofWeek  string `gorm:"size:32"`
	Location   string `gorm:"size:1024"`
	EntryId    int64  `gorm:"size:32;index"`
	CreatedAt  int64
	UpdatedAt  int64
}

// 域名资源

type Domain struct {
	Id          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"index"`
	VendorId    uint   `gorm:"index;default:null"`
	Name        string `gorm:"size:255"`
	NSList      string `gorm:"size:1024"`
	Model       string `gorm:"size:32"`
	CloudId     string `gorm:"size:64;uniqueIndex"`
	CloudMeta   any    `gorm:"serializer:json"`
	Status      string `gorm:"size:32"`
	Description string `gorm:"size:2048"`
	CreatedAt   int64
	UpdatedAt   int64
}

// 密钥对

type Keypair struct {
	Id          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"index"`
	PublicKey   string `gorm:"size:1024"`
	PrivateKey  string `gorm:"type:text" json:"-"`
	KeyType     string `gorm:"size:32;index"`
	Cipher      string `gorm:"size:64"`
	Status      string `gorm:"size:32"`
	Description string `gorm:"size:2048"`
	CreatedAt   int64
	UpdatedAt   int64
}

// 主机资源

type Machine struct {
	Id          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"index"`
	VendorId    uint   `gorm:"index;default:null"`
	HostName    string `gorm:"size:255"`
	IpAddress   string `gorm:"size:1024"`
	OSType      string `gorm:"size:32"`
	Region      string `gorm:"size:64"`
	Model       string `gorm:"size:32"`
	CloudId     string `gorm:"size:64;uniqueIndex;default:null"`
	CloudMeta   any    `gorm:"serializer:json"`
	WorkerId    string `gorm:"size:64;uniqueIndex;default:null"`
	WorkerMeta  any    `gorm:"serializer:json"`
	Status      string `gorm:"size:32"`
	Description string `gorm:"size:2048"`
	CreatedAt   int64
	UpdatedAt   int64
}

// 迁移记录

type Migration struct {
	Id          uint   `gorm:"primaryKey"`
	Version     string `gorm:"size:64;uniqueIndex"`
	Description string `gorm:"size:2048"`
	CreatedAt   int64
	UpdatedAt   int64
}

// 命令脚本

type Script struct {
	Id            uint   `gorm:"primaryKey"`
	UserId        uint   `gorm:"index"`
	Name          string `gorm:"size:128"`
	CommandType   string `gorm:"size:32"`
	Username      string `gorm:"size:64"`
	WorkDirectory string `gorm:"size:256"`
	Content       string `gorm:"type:text"`
	Timeout       uint
	Description   string `gorm:"size:2048"`
	CreatedAt     int64
	UpdatedAt     int64
}

// 任务记录

type Taskline struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index"`
	Subject   string `gorm:"size:128"`
	HostName  string `gorm:"size:128"`
	WorkerId  string `gorm:"size:64;index"`
	Request   any    `gorm:"serializer:json"`
	Response  any    `gorm:"serializer:json"`
	Status    string `gorm:"size:32"`
	CreatedAt int64
	UpdatedAt int64
}

// 用户

type User struct {
	Id          uint     `gorm:"primaryKey"`
	Username    string   `gorm:"size:64;uniqueIndex"`
	Password    string   `gorm:"size:128" json:"-"`
	Level       uint     `gorm:"default:5"`
	AppId       string   `gorm:"size:128;uniqueIndex"`
	AppKey      string   `gorm:"size:128" json:"-"`
	Email       string   `gorm:"size:255;uniqueIndex"`
	Avatar      string   `gorm:"size:255;default:assets/image/avatar.jpg"`
	Description string   `gorm:"size:2048;default:挥一挥手"`
	Vendors     []Vendor `json:",omitempty"`
	CreatedAt   int64
	UpdatedAt   int64
}

// 厂商

type Vendor struct {
	Id          uint      `gorm:"primaryKey"`
	UserId      uint      `gorm:"index"`
	SecretId    string    `gorm:"size:128;uniqueIndex"`
	SecretKey   string    `gorm:"size:128" json:"-"`
	Provider    string    `gorm:"size:32"`
	Cipher      string    `gorm:"size:64"`
	Status      string    `gorm:"size:32"`
	Description string    `gorm:"size:2048"`
	Certjobs    []Certjob `json:",omitempty"`
	Domains     []Domain  `json:",omitempty"`
	Machines    []Machine `json:",omitempty"`
	CreatedAt   int64
	UpdatedAt   int64
}
