package server

import (
	"flag"

	"tdp-cloud/cmd"
)

var (
	Dsn    string
	Listen string
)

func Flags() *cmd.FlagSet {

	cmd := &cmd.FlagSet{
		FlagSet: flag.NewFlagSet("server", flag.ExitOnError),
		Comment: "服务端",
	}

	cmd.StringVar(&Dsn, "dsn", "cloud.db", "数据源名称，支持MySQL和SQLite")
	cmd.StringVar(&Listen, "listen", ":7800", "服务端监听的IP地址和端口")

	return cmd

}
