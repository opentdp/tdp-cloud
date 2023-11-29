package cmd

import (
	"fmt"
	"os"

	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/subset"
)

func Execute() {

	// 设置子命令集

	flagsets := subset.NewFlagSets()

	// 尝试解析子命令

	if len(os.Args) > 1 {
		if sub := flagsets[os.Args[1]]; sub != nil {
			sub.Parse(os.Args[2:])
			sub.Execute()
			return
		}
	}

	// 显示帮助信息

	fmt.Printf("%s\n%s\n\n", args.AppName, args.AppSummary)
	for k, v := range flagsets {
		fmt.Printf("[%s] %s\n\n", k, v.Comment)
		v.PrintDefaults()
		fmt.Println()
	}
	fmt.Println(args.ReadmeText)

	os.Exit(0)

}
