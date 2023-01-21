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

	cmd := &cmd.FlagSet{
		FlagSet: flag.NewFlagSet("service", flag.ExitOnError),
		Comment: "服务管理",
	}

	cmd.StringVar(&vInstall, "install", "", "安装服务 [server|worker]")
	cmd.StringVar(&vUninstall, "uninstall", "", "卸载服务 [server|worker]")

	return cmd

}
