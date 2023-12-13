package args

import (
	"embed"
)

// 调试模式

var Debug bool

// 嵌入文件

var Efs *embed.FS

// 数据存储

var Dataset = struct {
	Dir    string
	Secret string
}{
	Dir: "var",
}

// 日志参数

var Logger = struct {
	Dir    string
	Level  string
	Target string
}{
	Level:  "info",
	Target: "stdout",
}

// 数据库参数 - Server

var Database = struct {
	Type   string
	Host   string
	User   string
	Passwd string
	Name   string
	Option string
}{
	Type: "sqlite",
	Name: "server.db",
}

// 主节点参数 - Server

var Server = struct {
	Listen string
	JwtKey string
}{
	Listen: ":7800",
}

// 子节点参数 - Worker

var Worker = struct {
	Remote string
}{}
