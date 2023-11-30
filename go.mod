module tdp-cloud

go 1.21

require (
	// 加解密
	github.com/forgoer/openssl v1.6.0
	// WEB 框架
	github.com/gin-gonic/gin v1.9.1
	// JWT 认证支持
	github.com/golang-jwt/jwt/v4 v4.5.0
	// 生成 UUID
	github.com/google/uuid v1.4.0
	// 服务管理
	github.com/kardianos/service v1.2.2
	// 配置解析
	github.com/knadh/koanf/parsers/yaml v0.1.0
	github.com/knadh/koanf/providers/file v0.1.0
	github.com/knadh/koanf/v2 v2.0.1
	// Map 转结构体
	github.com/mitchellh/mapstructure v1.5.0
	// 辅助类库
	github.com/opentdp/go-helper v0.4.2
	// 计划任务
	github.com/robfig/cron/v3 v3.0.1
	// Go 官方模块
	golang.org/x/crypto v0.15.0
	golang.org/x/net v0.18.0
)

// 间接依赖包

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.4 // indirect
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.0.5 // indirect
	github.com/alibabacloud-go/debug v1.0.0 // indirect
	github.com/alibabacloud-go/openapi-util v0.1.0 // indirect
	github.com/alibabacloud-go/tea v1.2.1 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.5 // indirect
	github.com/alibabacloud-go/tea-utils/v2 v2.0.4 // indirect
	github.com/alibabacloud-go/tea-xml v1.1.3 // indirect
	github.com/aliyun/credentials-go v1.3.2 // indirect
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/caddyserver/certmagic v0.19.2 // indirect
	github.com/cheggaaa/pb/v3 v3.1.4 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/glebarez/go-sqlite v1.21.2 // indirect
	github.com/glebarez/sqlite v1.10.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.16.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/libdns/alidns v1.0.3 // indirect
	github.com/libdns/cloudflare v0.1.0 // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/libdns/tencentcloud v1.0.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20231016141302-07b5767bb0ed // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mholt/acmez v1.2.0 // indirect
	github.com/miekg/dns v1.1.57 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/shirou/gopsutil/v3 v3.23.10 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.803 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/term v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	gorm.io/gorm v1.25.5 // indirect
	modernc.org/libc v1.34.11 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.7.2 // indirect
	modernc.org/sqlite v1.27.0 // indirect
)
