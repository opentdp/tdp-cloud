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

	flag.StringVar(&Dsn, "dsn", "cloud.db", "数据来源名称")
	flag.StringVar(&Address, "address", ":7800", "服务监听地址和端口")

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
