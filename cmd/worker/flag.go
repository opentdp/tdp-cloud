package worker

import (
	"flag"

	"tdp-cloud/cmd"
)

var (
	vRemote string
)

func Flags() *cmd.FlagSet {

	command := &cmd.FlagSet{
		FlagSet: flag.NewFlagSet("worker", flag.ExitOnError),
		Comment: "客户端",
	}

	command.StringVar(&vRemote, "remote", "", `客户端注册地址（e.g. "ws://{domain}/wsi/{appid}/worker"）`)

	return command

}
