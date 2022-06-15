package cli

import (
	"flag"
	"fmt"
	"os"
)

var (
	Dsn     string
	Address string
)

func Flags() {

	flag.StringVar(&Dsn, "dsn", "cloud.db", "数据库存储路径")

	flag.StringVar(&Address, "address", ":7800", "服务器监听地址和端口")

	flag.Usage = usage

	flag.Parse()
}

func usage() {

	fmt.Fprintf(os.Stderr, `轻量服务器控制面板，项目地址 https://github.com/tdp-resource/tdp-cloud

可选参数:

`)

	flag.PrintDefaults()
}
