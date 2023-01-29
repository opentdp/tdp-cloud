module tdp-cloud

go 1.18

// WEB 框架

require (
	github.com/gin-gonic/gin v1.8.2
	github.com/gorilla/websocket v1.5.0
)

// 数据库 ORM

require (
	github.com/glebarez/sqlite v1.6.0
	gorm.io/driver/mysql v1.4.5
	gorm.io/gorm v1.24.3
)

// 工具类

require (
	// 生成 UUID
	github.com/google/uuid v1.3.0
	// HTTP 客户端
	github.com/imroc/req/v3 v3.31.0
	// 服务管理
	github.com/kardianos/service v1.2.2
	// Map 转结构体
	github.com/mitchellh/mapstructure v1.5.0
	// 计划任务
	github.com/robfig/cron/v3 v3.0.1
	// 系统信息采集
	github.com/shirou/gopsutil/v3 v3.22.12
	// 类型转换
	github.com/spf13/cast v1.5.0
	// 腾讯云 SDK
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.582
	// 加密扩充库
	golang.org/x/crypto v0.5.0
	// 终端和控制台支持包
	golang.org/x/term v0.4.0
	// 文本操作
	golang.org/x/text v0.6.0
)

// 间接依赖

require (
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/glebarez/go-sqlite v1.20.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20230110061619-bbe2e5e100de // indirect
	github.com/marten-seemann/qpack v0.2.1 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.3 // indirect
	github.com/marten-seemann/qtls-go1-19 v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230126093431-47fa9a501578 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/ugorji/go/codec v1.2.8 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	modernc.org/libc v1.22.2 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.20.3 // indirect
)
