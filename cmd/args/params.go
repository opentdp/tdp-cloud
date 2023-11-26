package args

import (
	"embed"

	"github.com/opentdp/go-helper/strutil"
)

// 调试模式

var Debug bool

// 嵌入文件

var Efs *embed.FS

// 数据存储

type IDataset struct {
	Dir    string
	Secret string
}

var Dataset = IDataset{
	Dir:    "var",
	Secret: strutil.Rand(32),
}

// 日志参数

type ILogger struct {
	Dir    string
	Level  string
	Target string
}

var Logger = ILogger{
	Dir:    ".",
	Level:  "info",
	Target: "stdout",
}

// 数据库参数

type IDatabase struct {
	Type   string
	Host   string
	User   string
	Passwd string
	Name   string
	Option string
}

var Database = IDatabase{
	Type: "sqlite",
	Name: "server.db",
}

// 主节点参数

type IServer struct {
	Listen string
	JwtKey string
}

var Server = IServer{
	Listen: ":7800",
	JwtKey: strutil.Rand(32),
}

// 子节点参数

type IWorker struct {
	Remote string
}

var Worker = IWorker{}
