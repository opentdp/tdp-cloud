package cli

import (
	"flag"
	"fmt"
	"os"
)

var (
	Address string
	Agent   string
	Dsn     string
)

func Flags() {

	// 服务器模式
	flag.StringVar(&Address, "address", ":7800", "服务端监听地址和端口")
	flag.StringVar(&Dsn, "dsn", "./cloud.db", "数据源名称，支持MySQL和SQLite")

	// 客户端模式
	flag.StringVar(&Agent, "agent", "", `注册为客户端（e.g. "ws://ip:7800/wsi/agent/*"）`)

	flag.Usage = usage

	flag.Parse()

}

func usage() {

	fmt.Fprintf(os.Stdout, `
轻量服务器控制面板

可选参数:

`)

	flag.PrintDefaults()

	fmt.Fprintf(os.Stdout, `
开源项目 https://github.com/tdp-resource/tdp-cloud
问题提交 https://github.com/tdp-resource/tdp-cloud/issues

`)

}
