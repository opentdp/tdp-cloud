package args

import (
	"fmt"
	"os"

	"tdp-cloud/cmd"
	"tdp-cloud/cmd/server"
	"tdp-cloud/cmd/service"
	"tdp-cloud/cmd/worker"
)

func Parser() {

	commands := getCommands()

	// 未输入子命令
	if len(os.Args) < 2 {
		showUsage(commands)
		return
	}

	// 设置全局子命令
	cmd.SubCommand = os.Args[1]

	// 尝试解析子命令
	if sub := commands[cmd.SubCommand]; sub != nil {
		sub.Parse(os.Args[2:])
	} else {
		showUsage(commands)
	}

}

func getCommands() cmd.FlagSets {

	se := server.Flags()
	wo := worker.Flags()

	sr := service.Flags()

	return cmd.FlagSets{
		se.Name(): se,
		wo.Name(): wo,
		sr.Name(): sr,
	}

}

func showUsage(commands cmd.FlagSets) {

	fmt.Println(header)

	for _, v := range commands {
		fmt.Printf("%s %s\n\n", v.Name(), v.Comment)
		v.PrintDefaults()
		fmt.Println()
	}

	fmt.Println(footer)

	os.Exit(2)

}
