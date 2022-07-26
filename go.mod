module tdp-cloud

go 1.18

// WEB 框架

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/gorilla/websocket v1.5.0
)

// 数据库 ORM

require (
	github.com/glebarez/sqlite v1.4.6
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

// 腾讯云 SDK

require (
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam v1.0.449
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.449
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod v1.0.449
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse v1.0.449
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor v1.0.449
)

// 其他依赖

require (
	// 计划任务
	github.com/robfig/cron/v3 v3.0.1
	// 日志输出
	github.com/sirupsen/logrus v1.9.0
	// 加密扩充库
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
)

// 间接依赖

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/glebarez/go-sqlite v1.17.3 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.10 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	modernc.org/libc v1.16.17 // indirect
	modernc.org/mathutil v1.4.1 // indirect
	modernc.org/memory v1.1.1 // indirect
	modernc.org/sqlite v1.17.3 // indirect
)
