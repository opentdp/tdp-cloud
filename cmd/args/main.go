package args

import (
	"fmt"
	"os"

	"tdp-cloud/cmd"
	"tdp-cloud/cmd/server"
	"tdp-cloud/cmd/worker"
)

var CmdName string

func Parser() {

	commands := getCommands()

	// 未输入子命令
	if len(os.Args) < 2 {
		showUsage(commands)
		return
	}

	// 设置全局参数
	CmdName = os.Args[1]

	// 尝试解析子命令
	if cmd := commands[CmdName]; cmd != nil {
		cmd.Parse(os.Args[2:])
		return
	}

	// 显示全局帮助
	showUsage(commands)

}

func getCommands() map[string]*cmd.FlagSet {

	s := server.Flags()
	w := worker.Flags()

	return map[string]*cmd.FlagSet{
		s.Name(): s,
		w.Name(): w,
	}

}

func showUsage(commands map[string]*cmd.FlagSet) {

	fmt.Println(header)

	for _, v := range commands {
		fmt.Printf("%s %s\n\n", v.Name(), v.Comment)
		v.PrintDefaults()
		fmt.Println()
	}

	fmt.Println(footer)

	os.Exit(2)

}
