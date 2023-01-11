package args

import (
	"flag"
	"fmt"
	"os"
)

var (
	Dsn    string
	Listen string
	Server string
)

func Flags() {

	// 服务器模式
	flag.StringVar(&Dsn, "dsn", "cloud.db", "数据源名称，支持MySQL和SQLite")
	flag.StringVar(&Listen, "listen", ":7800", "服务端监听的IP地址和端口")

	// 客户端模式
	flag.StringVar(&Server, "server", "", `客户端注册地址（e.g. "ws://ip:7800/wsi/*/worker"）`)

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
