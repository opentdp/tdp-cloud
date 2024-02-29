package args

import (
	"embed"
)

// 调试模式

var Debug bool

// 嵌入文件

var Efs *embed.FS

// 数据存储

type IDataset struct {
	Dir    string `yaml:"dir"`
	Secret string `yaml:"secret"`
}

var Dataset = &IDataset{
	Dir: "var",
}

// 日志参数

type ILogger struct {
	Dir    string `yaml:"dir"`
	Level  string `yaml:"level"`
	Target string `yaml:"target"`
}

var Logger = &ILogger{
	Level:  "info",
	Target: "stdout",
}

// 数据库参数 - Server

type IDatabase struct {
	Type   string `yaml:"type"`
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Name   string `yaml:"name"`
	Option string `yaml:"option"`
}

var Database = &IDatabase{
	Type: "sqlite",
	Name: "server.db",
}

// 主节点参数 - Server

type IServer struct {
	Listen string `yaml:"listen"`
	JwtKey string `yaml:"jwtkey"`
}

var Server = &IServer{
	Listen: ":7800",
}

// 子节点参数 - Worker

type IWorker struct {
	Remote string `yaml:"remote"`
}

var Worker = &IWorker{}
