package worker

import (
	"flag"

	"tdp-cloud/cmd"
)

var (
	vRemote string
)

func Flags() *cmd.FlagSet {

	cmd := &cmd.FlagSet{
		FlagSet: flag.NewFlagSet("worker", flag.ExitOnError),
		Comment: "客户端",
	}

	cmd.StringVar(&vRemote, "remote", "", `客户端注册地址（e.g. "ws://{domain}/wsi/{appid}/worker"）`)

	return cmd

}
