package dborm

// 用户表

type User struct {
	Id          uint   `gorm:"primaryKey"`
	Username    string `gorm:"index,unique"`
	Password    string
	Description string `gorm:"default:不可能！我的代码怎么可能会有bug！"`
	Secrets     []Secret
	Sessions    []Session
	CreatedAt   int64
	UpdatedAt   int64
}

// 密钥表

type Secret struct {
	Id          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"index"`
	SecretId    string `gorm:"index,unique"`
	SecretKey   string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 会话表

type Session struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index"`
	Token     string `gorm:"index,unique"`
	CreatedAt int64
	UpdatedAt int64
}

// 自动化助手

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

type TATHistory struct {
	Id                   uint `gorm:"primaryKey"`
	UserId               uint `gorm:"index"`
	KeyId                uint
	Name                 string
	Region               string
	InvocationId         string
	InvocationStatus     string
	InvocationResultJson string
}
