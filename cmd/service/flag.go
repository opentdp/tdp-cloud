package service

import (
	"flag"

	"tdp-cloud/cmd"
)

var (
	vInstall   string
	vUninstall string
)

func Flags() *cmd.FlagSet {

	command := &cmd.FlagSet{
		FlagSet: flag.NewFlagSet("service", flag.ExitOnError),
		Comment: "服务管理",
	}

	command.StringVar(&vInstall, "install", "", "安装服务 [ server | worker ] [ - 其他参数 ]")
	command.StringVar(&vUninstall, "uninstall", "", "卸载服务 [ server | worker ]")

	return command

}
