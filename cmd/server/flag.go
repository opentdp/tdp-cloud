package server

import (
	"flag"

	"tdp-cloud/cmd"
)

var (
	vDsn    string
	vListen string
)

func Flags() *cmd.FlagSet {

	command := &cmd.FlagSet{
		FlagSet: flag.NewFlagSet("server", flag.ExitOnError),
		Comment: "服务端",
	}

	command.StringVar(&vDsn, "dsn", "server.db", "数据源名称，支持MySQL和SQLite")
	command.StringVar(&vListen, "listen", ":7800", "服务端监听的IP地址和端口")

	return command

}
