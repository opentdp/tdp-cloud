package cmd

import (
	"fmt"
	"os"

	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/subset"
)

func Execute() {

	// 设置子命令集

	commands := subset.NewFlagSets()

	// 尝试解析子命令

	if len(os.Args) > 1 {
		if cmd := commands[os.Args[1]]; cmd != nil {
			cmd.Parse(os.Args[2:])
			cmd.Execute()
			return
		}
	}

	// 显示帮助信息

	fmt.Printf("%s\n%s\n\n", args.AppName, args.AppSummary)
	for k, v := range commands {
		fmt.Printf("[%s] %s\n\n", k, v.Comment)
		v.PrintDefaults()
		fmt.Println()
	}
	fmt.Println(args.ReadmeText)

	os.Exit(0)

}
